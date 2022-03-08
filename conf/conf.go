package conf

import (
	"es-entertainment/core/database/mysql"

	"es-entertainment/core/database/redis"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	}

	Redis redis.RedisStruct
	Mysql map[string]mysql.MysqlStruct
}

func GetConf() Config {
	var cfg Config
	v := viper.New()
	v.SetConfigName("app")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := v.Unmarshal(&cfg); err != nil {
		panic(err)
	}
	return cfg
}
