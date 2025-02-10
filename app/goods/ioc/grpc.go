package ioc

import (
	"Toktik/app/goods/conf"
	"Toktik/app/goods/idl/goods"
	"Toktik/app/goods/server"
	"Toktik/app/goods/utils/grpcs"
	"context"
	"strconv"

	"google.golang.org/grpc"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func InitGRPCxServer(ctx context.Context, gd *server.GoodsServer, ecli *clientv3.Client) *grpcs.Server {
	configs := ctx.Value("config").(*conf.Config)
	port, err := strconv.Atoi(configs.ServerInfoConfig.Port)
	if err != nil {
		panic("端口错误")
	}
	servers := grpc.NewServer()
	goods.RegisterGoodsServiceServer(servers, gd)
	return &grpcs.Server{
		Server:     servers,
		Name:       "goods",
		Port:       port,
		EtcdClient: ecli,
		EtcdTTL:    int64(configs.ServerInfoConfig.EtcdTTL),
	}
}
