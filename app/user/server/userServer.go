package server

import (
	"Toktik-E-commerce/app/user/biz/service"
	users "Toktik-E-commerce/app/user/idl/userV"
	"context"
	"gorm.io/gorm"
)

type UserServer struct {
	ctx context.Context
	db  *gorm.DB
	rs  *service.RegisterService
	ls  *service.LoginService
	users.UnimplementedUserServiceServer
}

func NewUserServer(ctx context.Context, db *gorm.DB) *UserServer {
	return &UserServer{db: db,
		rs: service.NewRegisterService(ctx, db),
		ls: service.NewLoginService(ctx, db),
	}
}
func (au *UserServer) Register(ctx context.Context, registerReq *users.RegisterReq) (resp *users.RegisterResp, err error) {
	resp, err = au.rs.Run(registerReq)
	return
}
func (au *UserServer) Login(ctx context.Context, loginReq *users.LoginReq) (resp *users.LoginResp, err error) {
	resp, err = au.ls.Run(loginReq)
	return
}
