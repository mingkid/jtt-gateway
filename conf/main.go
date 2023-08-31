package conf

import (
	"fmt"

	"github.com/spf13/viper"
)

var (
	Database database
	Web      web
	JTT      jtt
)

type Config struct {
	Database database `yaml:"database"` // 数据库配置
	Web      web      `yaml:"web"`      // Web 服务配置
	JTT      jtt      `yaml:"jtt"`      // JTT 服务配置
}

type database struct {
	Host                  string `yaml:"host"`
	Port                  int    `yaml:"port"`
	User                  string `yaml:"user"`
	Password              string `yaml:"password"`
	Dbname                string `yaml:"dbname"`
	MaxIdleConns          int    `yaml:"max_idle_conns"`
	MaxOpenConns          int    `yaml:"max_open_conns"`
	ConnectionMaxLifetime int    `yaml:"connection_max_lifetime"`
}

type web struct {
	MaxPageSize uint `yaml:"max_page_size"`
	Port        uint `yaml:"port"`
}

type jtt struct {
	Port uint `yaml:"port"`
}

func init() {
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

	Database = conf.Database
	Web = conf.Web
	JTT = conf.JTT
}
