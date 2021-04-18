package cache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

var rds *redis.Client

func initClient() (err error) {
	rds = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:8377",
		Password: "20101269",
		DB:       0,   // 使用默认的DB （16个 0~15）
		PoolSize: 100, // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := rds.Ping(ctx).Result()
	fmt.Println("redis Ping()", res)
	return
}

func NewRedis() *redis.Client {
	err := initClient()
	if err != nil {
		log.Println("redis 连接失败")
		return nil
	}
	return rds
}
