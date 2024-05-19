package common

import (
	"GinVue/model"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

// DB 全局DB实例
var DB *gorm.DB

// InitDB 连接数据库
func InitDB() *gorm.DB {

	//一些配置信息，
	driverName := viper.GetString("datasource.driverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")

	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username, password, host, port, database, charset)

	//连接mysql数据库
	db, err := gorm.Open(driverName, args)

	if err != nil {
		panic("failed to connect databases,err:" + err.Error())
	}

	//数据库迁移字段
	db.AutoMigrate(&model.User{}, &model.Todo{})

	//外键关系
	db.Model(&model.Todo{}).AddForeignKey("telephone", "users(telephone)", "CASCADE", "CASCADE")
	DB = db
	return db
}

// GetDB 暴露的接口
func GetDB() *gorm.DB {
	return DB
}
