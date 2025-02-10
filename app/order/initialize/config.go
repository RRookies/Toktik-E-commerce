package initialize

import(
	"github.com/spf13/viper"
	"Tiktok/app/order/global"
)

func InitializeConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../")

	if err := viper.ReadInConfig(); err != nil{
		panic(err)
	}
	if err  := viper.Unmarshal(&global.ServerConfig); err != nil{
		panic(err)
	}
}