package model

import (
	"github.com/jinzhu/gorm"
)

func migration(DB *gorm.DB) {
	// 自动迁移模式
	DB.Set("gorm:table_options", "charset=utf8mb4")
	DB.AutoMigrate(&User{}).
		AutoMigrate(&Folder{}).
		AutoMigrate(&Article{}).
		AutoMigrate(&TimeLine{}).
		AutoMigrate(&MyBook{}).
		AutoMigrate(&MyMood{})
}
