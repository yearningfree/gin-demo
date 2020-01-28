package cache

import (
	"github.com/go-redis/redis"
)

// RedisCli redis缓存客户端单例
var RedisCli *redis.Client

// InitRedis 初始化Redis
func InitRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456",
		DB:       0, // use default DB
	})

	_, err := client.Ping().Result()

	if err != nil {
		panic(err)
	}

	RedisCli = client
}



