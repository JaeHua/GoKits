package main

import (
	"GinVue/common"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"os"
)

func main() {

	//配置初始化，配置文件路径等
	InitConfig()

	//初始化DB
	db := common.InitDB()

	//初始化RD
	rd := common.InitSingleRedis()

	//数据库延时关闭
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	//Redis延时关闭
	defer func(Db *redis.Client) {
		err := Db.Close()
		if err != nil {

		}
	}(rd.Db)

	//默认处理引擎
	r := gin.Default()

	//获取路由
	r = CollectRoute(r)

	//获得mysql服务端口
	port := viper.GetString("server.port")

	//监听端口
	panic(r.Run(":" + port))
}

// InitConfig 初始化配置
func InitConfig() {
	//工作目录
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
