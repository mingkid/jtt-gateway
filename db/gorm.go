package db

import (
	"fmt"
	"github.com/mingkid/jtt808-gateway/conf"
	"github.com/mingkid/jtt808-gateway/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	var err error

	config := conf.LoadConfig()
	DB, err = gorm.Open(postgres.Open(
		fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			config.Database.Host, config.Database.Port, config.Database.User, config.Database.Password, config.Database.Dbname),
	), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("数据库连接失败")
	}
	err = DB.AutoMigrate(model.Term{}, model.Platform{})
	if err != nil {
		fmt.Println(err.Error())
		panic("数据库迁移")
	}
}
