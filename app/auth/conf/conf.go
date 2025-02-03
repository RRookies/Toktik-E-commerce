package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"log/slog"
	"os"
	"sync"
)

type Config struct {
	Redis            Redis            `mapstructure:"redis"`
	JwtConfig        JwtConfig        `mapstructure:"jwt"`
	Registry         Registry         `mapstructure:"registry"`
	ServerInfoConfig ServerInfoConfig `mapstructure:"server_info"`
	EtcdInfo         EtcdInfo         `mapstructure:"etcd"`
	once             sync.Once
}

type Redis struct {
	Address  string `mapstructure:"address"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type JwtConfig struct {
	Key string `mapstructure:"key"`
}
type Registry struct {
	RegistryAddress string `mapstructure:"registry_address"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
}

type ServerInfoConfig struct {
	Port    string `mapstructure:"port"`
	EtcdTTL int    `mapstructure:"etcdttl"`
}

type EtcdInfo struct {
	Endpoints []string `mapstructure:"endpoints"`
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
