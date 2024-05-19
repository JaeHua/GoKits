package common

import (
	"GinVue/model"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"log"
)

var RD *model.RedisSingleObj

// InitSingleRedis 初始化redis数据库
func InitSingleRedis() (rd *model.RedisSingleObj) {

	//一些配置信息
	log.Println("hi")
	host := viper.GetString("redis.host")
	log.Println("rd.RedisHost", host)

	port := viper.GetString("redis.port")
	log.Println("rd.RedisPort", port)

	auth := viper.GetString("redis.auth")
	log.Println("rd.RedisAuth", auth)

	//配置实例

	rd = &model.RedisSingleObj{
		RedisHost: host,
		RedisPort: port,
		RedisAuth: auth,
	}
	//Redis连接格式拼接
	redisAddr := fmt.Sprintf("%s:%s", rd.RedisHost, rd.RedisPort)
	// Redis 连接对象: NewClient将客户端返回到由选项指定的Redis服务器
	rd.Db = redis.NewClient(&redis.Options{
		Addr:        redisAddr,    // redis服务ip:port
		Password:    rd.RedisAuth, // redis的认证密码
		DB:          rd.Database,  // 连接的database库
		IdleTimeout: 300,          // 默认Idle超时时间
		PoolSize:    100,          // 连接池
	})
	// 验证是否连接到redis服务端
	res, err := rd.Db.Ping().Result()
	RD = rd
	if err != nil {
		fmt.Printf("Connect Failed! Err: %v\n", err)
		return nil
	} else {
		fmt.Printf("Connect Successful! Ping => %v\n", res)
		return rd
	}
}

// GetRedis 暴露调用接口
func GetRedis() *model.RedisSingleObj {
	return RD
}
