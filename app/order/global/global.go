package global

import (
	"Tiktok/app/order/config"
	"Tiktok/app/order/idl/gen"
	"gorm.io/gorm"
	"go.etcd.io/etcd/client/v3"
)

var (
	DB *gorm.DB
	EtcdClient   *clientv3.Client
	ServerConfig config.ServerConfig
	ProductsSevClient gen.ProductCatalogServiceClient
)