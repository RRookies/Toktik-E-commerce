package service

import (
	"Toktik-E-commerce/app/auth/biz/model"
	auths "Toktik-E-commerce/app/auth/idl/authV1"
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/redis/go-redis/v9"
	"time"
)

type RefreshTokenService struct {
	ctx    context.Context
	client *redis.Client
	jwtKey []byte
}

func NewRefreshTokenService(ctx context.Context, client *redis.Client, jwtKey []byte) *RefreshTokenService {
	return &RefreshTokenService{
		ctx:    ctx,
		client: client,
		jwtKey: jwtKey,
	}
}

func (rts *RefreshTokenService) Run(ctx context.Context, req *auths.RefreshTokenReq) (*auths.RefreshResp, error) {
	exists, err := rts.client.Exists(ctx, req.Token).Result()
	if err != nil {
		return nil, err
	}
	if exists > 0 {
		return nil, err
	}
	token, err := jwt.ParseWithClaims(req.Token, &model.UserClaim{}, func(token *jwt.Token) (interface{}, error) {
		return rts.jwtKey, nil
	}, jwt.WithValidMethods([]string{"HS256"}))
	if err != nil {
		return nil, err
	}
	var userId int
	if claims, ok := token.Claims.(*model.UserClaim); ok && token.Valid {
		userId = claims.UserId
		claimsnew := model.UserClaim{
			UserId: int(userId),
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				NotBefore: jwt.NewNumericDate(time.Now()),
				Issuer:    "Server",
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsnew)
		signedString, err := token.SignedString(rts.jwtKey)
		if err != nil {
			return nil, err
		}
		return &auths.RefreshResp{
			Token: signedString,
		}, nil
	}
	return nil, err
}
