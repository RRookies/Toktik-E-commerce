package ioc

import (
	"Toktik-E-commerce/app/user/conf"
	users "Toktik-E-commerce/app/user/idl/userV"
	"Toktik-E-commerce/app/user/server"
	"Toktik-E-commerce/app/user/utils/grpcx"
	"context"
	"google.golang.org/grpc"
	"strconv"
)
import (
	clientv3 "go.etcd.io/etcd/client/v3"
)

func InitGRPCxServer(ctx context.Context, au *server.UserServer, ecli *clientv3.Client) *grpcx.Server {
	configs := ctx.Value("config").(*conf.Config)
	port, err := strconv.Atoi(configs.ServerInfoConfig.Port)
	if err != nil {
		panic("端口错误")
	}
	servers := grpc.NewServer()
	users.RegisterUserServiceServer(servers, au)
	return &grpcx.Server{
		Server:     servers,
		Name:       "user",
		Port:       port,
		EtcdClient: ecli,
		EtcdTTL:    int64(configs.ServerInfoConfig.EtcdTTL),
	}
}
