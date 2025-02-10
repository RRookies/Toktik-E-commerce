package config

type ProductsSrvConfig struct {
    Name string `mapstructure:"name" json:"name"`
}

type MysqlConfig struct {
    Host     string `mapstructure:"host" json:"host"`
    Port     int    `mapstructure:"port" json:"port"`
    Name     string `mapstructure:"db" json:"db"`
    User     string `mapstructure:"user" json:"user"`
    Password string `mapstructure:"password" json:"password"`
}

type EtcdConfig struct {
    Endpoints   []string `mapstructure:"endpoints" json:"endpoints"` // ETCD集群地址
    DialTimeout int      `mapstructure:"dial-timeout" json:"dial-timeout"` // 连接超时时间（秒）
    TTL         int      `mapstructure:"ttl" json:"ttl"` // 租约TTL（秒）
}


type ServerConfig struct {
    Name       string       `mapstructure:"name" json:"name"`
    Host       string       `mapstructure:"host" json:"host"`
    Tags       []string     `mapstructure:"tags" json:"tags"`
    MysqlInfo  MysqlConfig  `mapstructure:"mysql" json:"mysql"`
    EtcdInfo   EtcdConfig   `mapstructure:"etcd" json:"etcd"` // 新增ETCD配置 

    ProductsSrvInfo     ProductsSrvConfig `mapstructure:"products_srv" json:"products_srv"`
}