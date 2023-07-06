package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Database Database `yaml:"database"`
}

type Database struct {
	Host                  string `yaml:"host"`
	Port                  int    `yaml:"port"`
	User                  string `yaml:"user"`
	Password              string `yaml:"password"`
	Dbname                string `yaml:"dbname"`
	MaxIdleConns          int    `yaml:"max_idle_conns"`
	MaxOpenConns          int    `yaml:"max_open_conns"`
	ConnectionMaxLifetime int    `yaml:"connection_max_lifetime"`
}

func LoadConfig() *Config {
	conf := &Config{}
	v := viper.New()
	v.SetConfigFile("./config.yaml")
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".") // 可以在工作目录当中查找配置
	err := v.ReadInConfig()
	if err != nil {
		fmt.Println("配置文件加载错误")
		panic(err)
	}
	err = v.Unmarshal(conf)
	if err != nil {
		fmt.Println("解析配置文件错误")
		panic(err)
	}
	return conf
}
