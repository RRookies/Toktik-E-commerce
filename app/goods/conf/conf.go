package conf

import (
	"fmt"
	"log/slog"
	"os"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	MysqlInfo        MysqlInfo        `mapstructure:"mysql"`
	EtcdInfo         EtcdInfo         `mapstructure:"etcd"`
	ServerInfoConfig ServerInfoConfig `mapstructure:"server_info"`
	once             sync.Once
}

type MysqlInfo struct {
	Dsn string `mapstructure:"dsn"`
}
type EtcdInfo struct {
	Endpoints []string `mapstructure:"endpoints"`
}
type Redis struct {
	Address  string `mapstructure:"address"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}
type ServerInfoConfig struct {
	Port    string `mapstructure:"port"`
	EtcdTTL int    `mapstructure:"etcdttl"`
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) GetConfig() *Config {
	err := c.InitConfig()
	if err != nil {
		panic(err)
	}
	return c
}

func (c *Config) InitConfig() (err error) {
	c.once.Do(func() {
		vp := viper.New()
		vp.SetConfigName("config")
		vp.AddConfigPath("./conf/config")
		vp.SetConfigType("yaml")
		fmt.Println(os.Getwd())
		if err = vp.ReadInConfig(); err != nil {
			slog.Error("无法读取配置")
			return
		}
		if err = vp.Unmarshal(&c); err != nil {
			slog.Error("无法解析配置")
			return
		}
		slog.Info("%+v", c)
	})
	return
}
