package main

import (
	"Toktik-E-commerce/app/auth/biz/dal/redis"
	"Toktik-E-commerce/app/auth/biz/service"
	"Toktik-E-commerce/app/auth/conf"
	"Toktik-E-commerce/app/auth/ioc"
	"Toktik-E-commerce/app/auth/server"
	"context"
)

//func Run(ctx context.Context) {
//	configs := ctx.Value("config").(*conf.Config)
//	fmt.Printf("configs: %+v\n", configs)
//	l, err := net.Listen("tcp", fmt.Sprintf(":%s", configs.ServerInfoConfig.Port))
//	if err != nil {
//		slog.Error("auth servier:Error listening on port 50051")
//	}
//	grpcServer := grpc.NewServer()
//	auths.RegisterAuthServiceServer(grpcServer, server.NewAuthServer(configs.JwtConfig.Key))
//	slog.Info("auth server:listening on port 50051")
//	if err := grpcServer.Serve(l); err != nil {
//		slog.Error("auth server: %v", err)
//	}
//}

func main() {
	configs := conf.NewConfig()
	err := configs.InitConfig()
	if err != nil {
		return
	}
	_ = configs.Registry.RegistryAddress
	ctx := context.WithValue(context.Background(), "config", configs)
	rediss := redis.NewRedisClient(ctx)
	etcdcli := ioc.InitEtcdClient(ctx)
	dts := service.NewDeliverTokenService(ctx, []byte(configs.JwtConfig.Key))
	lts := service.NewLogoutTokenService(ctx, rediss)
	rts := service.NewRefreshTokenService(ctx, rediss, []byte(configs.JwtConfig.Key))
	vts := service.NewVerifyTokenService(ctx, rediss, []byte(configs.JwtConfig.Key))
	au := server.NewAuthServer(dts, lts, rts, vts)
	grpcServer := ioc.InitGRPCxServer(ctx, au, etcdcli)
	err = grpcServer.Serve()
	if err != nil {
		panic(err)
	}
}
