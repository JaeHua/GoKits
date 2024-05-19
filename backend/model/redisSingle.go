package model

import "github.com/go-redis/redis"

// RedisSingleObj 定义一个redis对象结构体
type RedisSingleObj struct {
	RedisHost string
	RedisPort string
	RedisAuth string
	Database  int
	Db        *redis.Client
}
