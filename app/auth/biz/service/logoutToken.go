package service

import (
	auths "Toktik-E-commerce/app/auth/idl/authV1"
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type LogoutTokenService struct {
	ctx    context.Context
	client *redis.Client
}

func NewLogoutTokenService(ctx context.Context, client *redis.Client) *LogoutTokenService {
	return &LogoutTokenService{ctx: ctx,
		client: client,
	}
}
func (service *LogoutTokenService) Run(ctx context.Context, req *auths.LogoutTokenReq) (resp *auths.LogoutTokenResp, err error) {
	err = service.client.SetNX(ctx, req.Token, "1", 24*time.Hour).Err()
	if err != nil {
		return &auths.LogoutTokenResp{Res: false}, err
	}
	return &auths.LogoutTokenResp{Res: true}, err
}
