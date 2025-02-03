package main

import (
	"Toktik-E-commerce/app/user/conf"
	"Toktik-E-commerce/app/user/ioc"
	"Toktik-E-commerce/app/user/server"
	"context"
)

func main() {
	configs := conf.NewConfig()
	err := configs.InitConfig()
	if err != nil {
		panic(err)
		return
	}
	ctx := context.WithValue(context.Background(), "config", configs)
	db := ioc.NewMysqlDB(ctx)
	etcdcli := ioc.InitEtcdClient(ctx)
	us := server.NewUserServer(ctx, db)
	grpcServer := ioc.InitGRPCxServer(ctx, us, etcdcli)
	err = grpcServer.Serve()
	if err != nil {
		panic(err)
	}
}
