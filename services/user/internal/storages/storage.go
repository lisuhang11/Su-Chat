package storages

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
	"time"
)

var (
	db      *gorm.DB
	once    sync.Once
	initErr error
)

func InitDB() (*gorm.DB, error) {
	once.Do(func() {
		dsn := "root:root@tcp(127.0.0.1:3306)/Su-Chat?charset=utf8mb4&parseTime=True&loc=Local"
		db, initErr = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if initErr != nil {
			return
		}
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)
	})
	return db, initErr
}
