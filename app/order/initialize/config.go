package initialize

import (
	"Tiktok/app/order/global"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func InitializeConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil{
		panic(err)
	}
	if err  := viper.Unmarshal(&global.ServerConfig); err != nil{
		panic(err)
	}
	zap.S().Info("配置信息加载成功")
}