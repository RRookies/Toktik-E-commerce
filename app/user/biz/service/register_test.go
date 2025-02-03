package service

import (
	"Toktik-E-commerce/app/user/idl/userV"
	"context"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log/slog"
	"testing"
)

func init_db() (db *gorm.DB, err error) {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("../../config")
	viper.ReadInConfig()
	dsn := viper.GetString("mysql.dsn")
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		slog.Error("db connect err:" + err.Error())
	}
	return
}
func TestRegister_run(t *testing.T) {
	db, err := init_db()
	if err != nil {
		t.Errorf("db init err:" + viper.GetString("mysql.dsn"))
	}
	rs := NewRegisterService(context.Background(), db)
	resp, err := rs.Run(&userV.RegisterReq{
		Email:    "996948441@qq.com",
		Username: "wenshuai",
		Nickname: "wenshuai1",
		Password: "123456789",
	})
	fmt.Println(resp)
	ls := NewLoginService(context.Background(), db)
	respp, err := ls.Run(&userV.LoginReq{
		Email:    "996948441@qq.com",
		Password: "123456789",
	})
	fmt.Println(respp)
}
