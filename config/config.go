package config

import (
	"MyBlog/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:root@tcp(127.0.0.1:3306)/myblog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("数据库连接失败：", err)
	}
	DB = db

	// 自动迁移（建表）
	err = DB.AutoMigrate(&models.Category{}, &models.Tag{}, &models.Article{})
	if err != nil {
		log.Fatal("自动迁移失败：", err)
	}
}
