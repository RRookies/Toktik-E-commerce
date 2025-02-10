package initialize

import (
	"Tiktok/app/order/global"
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
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

func RegisterService(serviceAddr string) (error) {
	lease,err := global.EtcdClient.Grant(context.Background(),global.ServerConfig.EtcdInfo.TTL)
	if err != nil{
		zap.S().Errorf("ETCD 注册租约失败: %v", err)
		return err
	}

	serviceKey := fmt.Sprintf("/services/%s/%s", global.ServerConfig.Name, serviceAddr)
	_,err = global.EtcdClient.Put(context.Background(),serviceKey,serviceAddr,clientv3.WithLease(lease.ID))
	if err!= nil{
		zap.S().Errorf("ETCD 注册失败: %v", err)
		return err
	}

	keepAliveChan,err := global.EtcdClient.KeepAlive(context.Background(),lease.ID)
	if err!= nil{
		zap.S().Errorf("ETCD 保持心跳失败: %v", err)
		return err
	}
	
	go func(){
		for{
			select {
			case _, ok := <-keepAliveChan:
				if !ok {
					zap.S().Errorf("ETCD 心跳失败: %v", err)
					return
				}
			}
		}
	}()

	return nil
}
