package Mydatabase

import (
	"github.com/go-redis/redis/v8"
	"time"
)

var client *redis.Client

func init() {
	InitRedisClient()
}

func InitRedisClient() {
	client = redis.NewClient(&redis.Options{
		Addr:        "47.108.158.210:6379", // 连接地址
		Password:    "123456",              // 密码
		DB:          0,                     // 数据库编号
		DialTimeout: 1 * time.Second,       // 链接超时
	})
}

func GetRedisClient() *redis.Client {
	return client
}
