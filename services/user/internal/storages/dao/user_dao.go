package dao

import (
	"Su-chat/services/user/internal/storages/models"
	"errors"
	"gorm.io/gorm"
)

// GetUserByUserID 根据用户ID查询用户
func GetUserByUserID(db *gorm.DB, userID string) (*models.User, error) {
	var user models.User
	err := db.Where("user_id = ?", userID).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}

// GetUserByPhone 根据手机号查询用户
func GetUserByPhone(db *gorm.DB, phone string) (*models.User, error) {
	var user models.User
	err := db.Where("phone = ?", phone).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}

// GetUserByEmail 根据邮箱查询用户
func GetUserByEmail(db *gorm.DB, email string) (*models.User, error) {
	var user models.User
	err := db.Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}

// CreateUser 插入新用户
func CreateUser(db *gorm.DB, user *models.User) error {
	return db.Create(user).Error
}

// UpdateUser 更新用户基本信息（支持部分字段）
func UpdateUser(db *gorm.DB, userID string, updates map[string]interface{}) error {
	return db.Model(&models.User{}).Where("user_id = ?", userID).Updates(updates).Error
}

// CheckPhoneEmailUnique 检查手机号和邮箱的唯一性（排除指定用户ID）
func CheckPhoneEmailUnique(db *gorm.DB, phone, email *string, excludeUserID string) error {
	if phone != nil && *phone != "" {
		var count int64
		query := db.Model(&models.User{}).Where("phone = ?", *phone)
		if excludeUserID != "" {
			query = query.Where("user_id != ?", excludeUserID)
		}
		if err := query.Count(&count).Error; err != nil {
			return err
		}
		if count > 0 {
			return errors.New("手机号已被使用")
		}
	}
	if email != nil && *email != "" {
		var count int64
		query := db.Model(&models.User{}).Where("email = ?", *email)
		if excludeUserID != "" {
			query = query.Where("user_id != ?", excludeUserID)
		}
		if err := query.Count(&count).Error; err != nil {
			return err
		}
		if count > 0 {
			return errors.New("邮箱已被使用")
		}
	}
	return nil
}
