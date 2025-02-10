package ioc

import (
	"Toktik/app/goods/biz/model"
	"Toktik/app/goods/conf"
	"context"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysqlDB(ctx context.Context) *gorm.DB {
	configs := ctx.Value("config").(*conf.Config)
	db, err := gorm.Open(mysql.Open(configs.MysqlInfo.Dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&model.Goods{})
	if err != nil {
		panic(err)
	}
	return db
}
