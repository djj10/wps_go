package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"wps_go/config"
	"wps_go/user_service/models"
)

var DB *gorm.DB

func InitDB() {
	config, err := config.LoadConfig("../config/config.yaml")
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	dsn := config.Database.Username + ":" + config.Database.Password + 
		"@tcp(" + config.Database.Host + ":" + fmt.Sprintf("%d", config.Database.Port) + 
		")/" + config.Database.Name + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 自动迁移
	db.AutoMigrate(&models.User{})

	DB = db
}