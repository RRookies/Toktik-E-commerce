package initialize

import (
	"Tiktok/app/order/global"
	"Tiktok/app/order/idl/gen"
	"context"
	"time"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitializeSrv() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	resp, err := global.EtcdClient.Get(ctx,global.ServerConfig.ProductsSrvInfo.Name)
	
	cancel()
	if err != nil {
		panic(err)
	}
	if len(resp.Kvs) == 0{
		panic("no products srv found")
	}
	serviceAddr := string(resp.Kvs[0].Value)

	conn,err := grpc.Dial(serviceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err!= nil {
		zap.S().Fatal("[InitSrvConn] 连接 【用户服务失败】")
	}

	global.ProductsSevClient = gen.NewProductCatalogServiceClient(conn)
}