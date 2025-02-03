package ioc

import (
	"Toktik-E-commerce/app/user/conf"
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func InitEtcdClient(ctx context.Context) *clientv3.Client {
	configs := ctx.Value("config").(*conf.Config)
	client, err := clientv3.New(clientv3.Config{
		Endpoints: configs.EtcdInfo.Endpoints,
	})
	if err != nil {
		panic(err)
	}
	return client
}
