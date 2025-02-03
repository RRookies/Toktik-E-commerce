package service

import (
	"Toktik-E-commerce/app/auth/biz/model"
	auths "Toktik-E-commerce/app/auth/idl/authV1"
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/redis/go-redis/v9"
)

type VerifyTokenService struct {
	ctx    context.Context
	client *redis.Client
	jwtKey []byte
}

func NewVerifyTokenService(ctx context.Context, client *redis.Client, jwtKey []byte) *VerifyTokenService {
	return &VerifyTokenService{
		ctx:    ctx,
		client: client,
		jwtKey: jwtKey,
	}
}

func (vts *VerifyTokenService) Run(ctx context.Context, req *auths.VerifyTokenReq) (*auths.VerifyResp, error) {
	exists, err := vts.client.Exists(ctx, req.Token).Result()
	if err != nil {
		return nil, err
	}
	if exists > 0 {
		return &auths.VerifyResp{Res: false}, err
	}
	token, err := jwt.ParseWithClaims(req.Token, &model.UserClaim{}, func(token *jwt.Token) (interface{}, error) {
		return vts.jwtKey, nil
	}, jwt.WithValidMethods([]string{"HS256"}))
	if err != nil {
		return &auths.VerifyResp{Res: false}, err
	}
	if _, ok := token.Claims.(*model.UserClaim); ok && token.Valid {
		return &auths.VerifyResp{Res: true}, nil
	}
	return &auths.VerifyResp{Res: false}, err

}
