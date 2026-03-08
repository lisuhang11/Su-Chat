package models

import (
	"time"
)

type User struct {
	ID           uint      `gorm:"primaryKey;autoIncrement"`
	UserType     int       `gorm:"column:user_type"`
	UserID       string    `gorm:"column:user_id;uniqueIndex:uniq_userid"`
	Nickname     string    `gorm:"column:nickname"`
	Password     string    `gorm:"column:password"`
	UserPortrait *string   `gorm:"column:user_portrait"`
	Phone        *string   `gorm:"uniqueIndex:uniq_phone"`
	Email        *string   `gorm:"uniqueIndex:uniq_email"`
	CreatedTime  time.Time `gorm:"column:created_time;autoCreateTime"`
	UpdatedTime  time.Time `gorm:"column:updated_time;autoUpdateTime"`
}

type UserExt struct {
	ID          uint   `gorm:"primaryKey"`
	UserID      string `gorm:"column:user_id;uniqueIndex:uniq_item_key"`
	ItemKey     string `gorm:"column:item_key;uniqueIndex:uniq_item_key"`
	ItemValue   string
	ItemType    int
	UpdatedTime time.Time `gorm:"autoUpdateTime"`
}
