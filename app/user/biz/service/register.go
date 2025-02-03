package service

import (
	"Toktik-E-commerce/app/user/biz/model"
	"Toktik-E-commerce/app/user/idl/userV"
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegisterService struct {
	db  *gorm.DB
	ctx context.Context
}

func NewRegisterService(ctx context.Context, db *gorm.DB) *RegisterService {
	return &RegisterService{ctx: ctx,
		db: db,
	}
}

func (s *RegisterService) Run(req *userV.RegisterReq) (resp *userV.RegisterResp, err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	newUser := &model.User{
		Email:    req.Email,
		Password: string(hashedPassword),
		Nickname: req.Nickname,
		Username: req.Username,
	}
	if err = model.Create(s.db, s.ctx, newUser); err != nil {
		return nil, errors.New("创建用户失败")
	}
	return &userV.RegisterResp{UserId: int32(newUser.ID)}, nil
}
