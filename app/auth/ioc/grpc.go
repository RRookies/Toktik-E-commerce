package ioc

import (
	"Toktik-E-commerce/app/auth/conf"
	auths "Toktik-E-commerce/app/auth/idl/authV1"
	"Toktik-E-commerce/app/auth/server"
	"Toktik-E-commerce/app/auth/utils/grpcx"
	"context"
	"google.golang.org/grpc"
	"strconv"
)
import (
	clientv3 "go.etcd.io/etcd/client/v3"
)

func InitGRPCxServer(ctx context.Context, au *server.AuthServer, ecli *clientv3.Client) *grpcx.Server {

	configs := ctx.Value("config").(*conf.Config)
	port, err := strconv.Atoi(configs.ServerInfoConfig.Port)
	if err != nil {
		panic("端口错误")
	}
	servers := grpc.NewServer()
	auths.RegisterAuthServiceServer(servers, au)
	return &grpcx.Server{
		Server:     servers,
		Name:       "auth",
		Port:       port,
		EtcdClient: ecli,
		EtcdTTL:    int64(configs.ServerInfoConfig.EtcdTTL),
	}
}
