package initialize

import (
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	"Tiktok/app/order/global"
	"time"
)

func InitETCD() {

	cfg := global.ServerConfig.EtcdInfo

	client, err := clientv3.New(clientv3.Config{
		Endpoints:   cfg.Endpoints,
		DialTimeout: time.Duration(cfg.DialTimeout) * time.Second,
	})
	if err != nil {
		panic(fmt.Sprintf("ETCD 初始化失败: %v", err))
	}

	global.EtcdClient = client
	zap.S().Info("ETCD 连接成功")
}