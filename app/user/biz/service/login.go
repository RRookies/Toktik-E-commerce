package service

import (
	"Toktik-E-commerce/app/user/biz/model"
	"Toktik-E-commerce/app/user/idl/userV"
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginService struct {
	db  *gorm.DB
	ctx context.Context
}

func NewLoginService(ctx context.Context, db *gorm.DB) *LoginService {
	return &LoginService{
		db:  db,
		ctx: ctx,
	}
}

func (ls *LoginService) Run(req *userV.LoginReq) (resp *userV.LoginResp, err error) {
	var user *model.User = nil
	if req.Email == "" && req.Username == "" {
		return nil, errors.New("邮箱或用户名不能为空")
	}
	if req.Email != "" && req.Username != "" {
		return nil, errors.New("系统错误")
	}
	if req.Email != "" {
		user, err = model.GetByEmail(ls.db, ls.ctx, req.Email)
	}
	if req.Username != "" {
		user, err = model.GetByUserName(ls.db, ls.ctx, req.Username)
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("密码或用户名错误")
	}
	return &userV.LoginResp{
		UserId: int32(user.ID),
	}, nil
}
