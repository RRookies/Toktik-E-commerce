package service

import (
	"Toktik-E-commerce/app/auth/biz/model"
	auths "Toktik-E-commerce/app/auth/idl/authV1"
	"context"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type DeliverTokenService struct {
	ctx    context.Context
	jwtKey []byte
}

func NewDeliverTokenService(ctx context.Context, jwtkey []byte) *DeliverTokenService {
	return &DeliverTokenService{
		ctx:    ctx,
		jwtKey: jwtkey,
	}
}

func (dts *DeliverTokenService) Run(ctx context.Context, req *auths.DeliverTokenReq) (resp *auths.DeliveryResp, error error) {
	claims := model.UserClaim{
		UserId: int(req.GetUserId()),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "Server",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString(dts.jwtKey)
	if err != nil {
		return nil, err
	}
	return &auths.DeliveryResp{
		Token: signedString,
	}, nil
}
