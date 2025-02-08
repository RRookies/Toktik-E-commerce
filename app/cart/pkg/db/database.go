package db

import (
	"Tiktok/app/cart/internal/model"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	dsn := viper.GetString("mysql.dsn")
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	// 自动迁移 CartItem 模型
	err = DB.AutoMigrate(&model.CartItem{})
	if err != nil {
		return err
	}
	return nil
}
