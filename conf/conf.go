package conf

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	}

	Redis struct {
		Host string
		Port int
		Auth string
	}
	Mysql map[string]struct {
		Host string
		Port int
		User string
		Pass string
		Db   string
	}
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
