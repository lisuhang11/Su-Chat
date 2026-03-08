package logic

import (
	"Su-chat/pkg/tools"
	"Su-chat/services/user/gen"
	"Su-chat/services/user/internal/storages/dao"
	"Su-chat/services/user/internal/storages/models"
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strings"
)

type userServer struct {
	gen.UnimplementedUserServiceServer
	db *gorm.DB
}

// NewUserServer 创建userServer实例
func NewUserServer(db *gorm.DB) *userServer {
	return &userServer{db: db}
}

// Register 用户注册
func (s *userServer) Register(ctx context.Context, req *gen.UserRegisterReq) (*gen.UserRegisterResp, error) {
	// 参数校验
	if req.Nickname == "" {
		return nil, errors.New("昵称不能为空")
	}
	if req.Password == "" {
		return nil, errors.New("密码不能为空")
	}
	// 手机号和邮箱至少提供一个
	if (req.Phone == nil || *req.Phone == "") && (req.Email == nil || *req.Email == "") {
		return nil, errors.New("手机号和邮箱不能同时为空")
	}

	// 检查手机号和邮箱唯一性
	if err := dao.CheckPhoneEmailUnique(s.db, req.Phone, req.Email, ""); err != nil {
		return nil, err
	}

	// 生成用户ID
	userID := uuid.New().String()

	// 密码加密
	hashedPwd, err := tools.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("密码加密失败")
	}

	// 创建用户对象
	user := &models.User{
		UserType:     0, // 默认普通用户
		UserID:       userID,
		Nickname:     req.Nickname,
		Password:     hashedPwd,
		UserPortrait: req.UserPortrait,
		Phone:        req.Phone,
		Email:        req.Email,
	}

	// 开启事务
	err = s.db.Transaction(func(tx *gorm.DB) error {
		// 保存用户基本信息
		if err := dao.CreateUser(tx, user); err != nil {
			return err
		}
		// 处理扩展字段
		if req.ExtFields != nil && len(req.ExtFields) > 0 {
			if err := dao.ReplaceUserExts(tx, userID, req.ExtFields); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &gen.UserRegisterResp{UserId: userID}, nil
}

// Login 用户登录（支持手机号/邮箱）
func (s *userServer) Login(ctx context.Context, req *gen.UserLoginReq) (*gen.UserLoginResp, error) {
	identifier := req.UserId // 登录标识（可能是手机号或邮箱）
	password := req.Password
	if identifier == "" || password == "" {
		return nil, errors.New("登录标识和密码不能为空")
	}

	var user *models.User
	var err error

	// 尝试按手机号查找
	if strings.Contains(identifier, "@") { // 简单判断是否邮箱
		user, err = dao.GetUserByEmail(s.db, identifier)
	} else {
		user, err = dao.GetUserByPhone(s.db, identifier)
	}
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("用户不存在")
	}

	// 验证密码
	if !tools.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("密码错误")
	}

	return &gen.UserLoginResp{UserId: user.UserID}, nil
}

// GetUserInfo 查询用户信息
func (s *userServer) GetUserInfo(ctx context.Context, req *gen.UserInfoReq) (*gen.UserInfoResp, error) {
	if req.UserId == "" {
		return nil, errors.New("用户ID不能为空")
	}

	user, err := dao.GetUserByUserID(s.db, req.UserId)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("用户不存在")
	}

	// 查询扩展字段
	exts, err := dao.GetUserExts(s.db, req.UserId)
	if err != nil {
		return nil, err
	}
	extFields := make(map[string]string)
	for _, ext := range exts {
		extFields[ext.ItemKey] = ext.ItemValue
	}

	// 组装account字段（优先返回手机号，否则返回邮箱）
	var account string
	if user.Phone != nil && *user.Phone != "" {
		account = *user.Phone
	} else if user.Email != nil && *user.Email != "" {
		account = *user.Email
	}

	resp := &gen.UserInfoResp{
		UserId:       user.UserID,
		Account:      account,
		Nickname:     user.Nickname,
		Phone:        user.Phone,
		Email:        user.Email,
		UserPortrait: user.UserPortrait,
		ExtFields:    extFields,
	}
	return resp, nil
}

// UpdateUserInfo 更新用户信息
func (s *userServer) UpdateUserInfo(ctx context.Context, req *gen.UserInfoUpdateReq) (*gen.UserInfoUpdateResp, error) {
	if req.UserId == "" {
		return nil, errors.New("用户ID不能为空")
	}

	// 检查用户是否存在
	existing, err := dao.GetUserByUserID(s.db, req.UserId)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, errors.New("用户不存在")
	}

	// 检查手机号和邮箱唯一性
	if err := dao.CheckPhoneEmailUnique(s.db, req.Phone, req.Email, req.UserId); err != nil {
		return nil, err
	}

	// 构建更新字段
	updates := make(map[string]interface{})
	if req.Nickname != nil {
		updates["nickname"] = *req.Nickname
	}
	if req.Phone != nil {
		updates["phone"] = *req.Phone
	}
	if req.Email != nil {
		updates["email"] = *req.Email
	}
	if req.UserPortrait != nil {
		updates["user_portrait"] = *req.UserPortrait
	}

	// 开启事务
	err = s.db.Transaction(func(tx *gorm.DB) error {
		// 更新基本信息
		if len(updates) > 0 {
			if err := dao.UpdateUser(tx, req.UserId, updates); err != nil {
				return err
			}
		}
		// 处理扩展字段（全量替换）
		if req.ExtFields != nil {
			if err := dao.ReplaceUserExts(tx, req.UserId, req.ExtFields); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	// 查询更新后的完整信息返回
	updatedUser, err := dao.GetUserByUserID(s.db, req.UserId)
	if err != nil {
		return nil, err
	}
	exts, _ := dao.GetUserExts(s.db, req.UserId)
	extFields := make(map[string]string)
	for _, ext := range exts {
		extFields[ext.ItemKey] = ext.ItemValue
	}
	account := ""
	if updatedUser.Phone != nil && *updatedUser.Phone != "" {
		account = *updatedUser.Phone
	} else if updatedUser.Email != nil && *updatedUser.Email != "" {
		account = *updatedUser.Email
	}
	resp := &gen.UserInfoUpdateResp{
		UserId:       updatedUser.UserID,
		Account:      account,
		Nickname:     updatedUser.Nickname,
		Phone:        updatedUser.Phone,
		Email:        updatedUser.Email,
		UserPortrait: updatedUser.UserPortrait,
		ExtFields:    extFields,
	}
	return resp, nil
}

// GetUsersOnlineStatus 批量查询用户在线状态（此处简单返回离线，实际应集成IM服务）
func (s *userServer) GetUsersOnlineStatus(ctx context.Context, req *gen.UsersOnlineStatusReq) (*gen.UsersOnlineStatusResp, error) {
	if len(req.UserIds) == 0 {
		return &gen.UsersOnlineStatusResp{Items: []*gen.UserOnlineStatus{}}, nil
	}
	items := make([]*gen.UserOnlineStatus, 0, len(req.UserIds))
	for _, uid := range req.UserIds {
		items = append(items, &gen.UserOnlineStatus{
			UserId:   uid,
			IsOnline: false, // 默认为离线，可根据需要集成实时状态服务
		})
	}
	return &gen.UsersOnlineStatusResp{Items: items}, nil
}
