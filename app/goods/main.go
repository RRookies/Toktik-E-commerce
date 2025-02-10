package main

import (
	"Toktik/app/goods/conf"
	"Toktik/app/goods/ioc"
	"Toktik/app/goods/server"
	"context"
)

func main() {
	configs := conf.NewConfig()
	err := configs.InitConfig()
	if err != nil {
		panic(err)
	}
	ctx := context.WithValue(context.Background(), "config", configs)
	db := ioc.NewMysqlDB(ctx)
	etcdcli := ioc.InitEtcdClient(ctx)
	gd := server.NewGoodsServer(ctx, db)
	grpcServer := ioc.InitGRPCxServer(ctx, gd, etcdcli)
	err = grpcServer.Serve()
	if err != nil {
		panic(err)
	}
}
