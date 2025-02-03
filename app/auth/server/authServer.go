package server

import (
	"Toktik-E-commerce/app/auth/biz/service"
	auths "Toktik-E-commerce/app/auth/idl/authV1"
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/redis/go-redis/v9"
)

type AuthServer struct {
	auths.UnimplementedAuthServiceServer
	client *redis.Client
	dts    *service.DeliverTokenService
	lts    *service.LogoutTokenService
	rts    *service.RefreshTokenService
	vts    *service.VerifyTokenService
}

type MyClaims struct {
	userId int `json:"user"`
	jwt.RegisteredClaims
}

func NewAuthServer(dts *service.DeliverTokenService, lts *service.LogoutTokenService, rts *service.RefreshTokenService, vts *service.VerifyTokenService) *AuthServer {
	return &AuthServer{
		dts: dts,
		lts: lts,
		rts: rts,
		vts: vts,
	}
}
func (au *AuthServer) DeliverTokenByRPC(ctx context.Context, req *auths.DeliverTokenReq) (*auths.DeliveryResp, error) {
	return au.dts.Run(ctx, req)
}

func (au *AuthServer) VerifyTokenByRPC(ctx context.Context, req *auths.VerifyTokenReq) (*auths.VerifyResp, error) {
	return au.vts.Run(ctx, req)
}
func (au *AuthServer) RefreshTokenByRPC(ctx context.Context, req *auths.RefreshTokenReq) (*auths.RefreshResp, error) {
	return au.rts.Run(ctx, req)
}

func (au *AuthServer) LogoutTokenByRPC(ctx context.Context, req *auths.LogoutTokenReq) (resp *auths.LogoutTokenResp, err error) {
	return au.lts.Run(ctx, req)
}
