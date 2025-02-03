package service

import (
	"Toktik-E-commerce/app/user/idl/userV"
	"context"
	"fmt"
	"github.com/spf13/viper"
	"testing"
)

func TestLogin_run(t *testing.T) {
	db, err := init_db()
	if err != nil {
		t.Errorf("db init err:" + viper.GetString("mysql.dsn"))
	}
	ls := NewLoginService(context.Background(), db)
	respp, err := ls.Run(&userV.LoginReq{
		Email:    "996948441@qq.com",
		Password: "123456789",
	})
	if err != nil {
		t.Errorf("login err:" + err.Error())
	}
	fmt.Println(respp)
}
