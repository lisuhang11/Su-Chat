package dao

import (
	"Su-chat/services/user/internal/storages/models"
	"gorm.io/gorm"
)

// GetUserExts 查询用户的所有扩展字段
func GetUserExts(db *gorm.DB, userID string) ([]models.UserExt, error) {
	var exts []models.UserExt
	err := db.Where("user_id = ?", userID).Find(&exts).Error
	return exts, err
}

// ReplaceUserExts 全量替换用户的扩展字段（删除旧的全部插入新的）
func ReplaceUserExts(db *gorm.DB, userID string, exts map[string]string) error {
	// 删除旧的扩展字段
	if err := db.Where("user_id = ?", userID).Delete(&models.UserExt{}).Error; err != nil {
		return err
	}
	// 插入新的扩展字段
	if len(exts) > 0 {
		var userExts []models.UserExt
		for key, value := range exts {
			userExts = append(userExts, models.UserExt{
				UserID:    userID,
				ItemKey:   key,
				ItemValue: value,
				ItemType:  0, // 默认类型
			})
		}
		return db.Create(&userExts).Error
	}
	return nil
}
