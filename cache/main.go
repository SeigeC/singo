package cache

import (
	"context"
	"os"
	"singo/util"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

type Client redis.Client

func (c Client) Context() (context.Context, func()) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}

// RedisClient Redis缓存客户端单例
var RedisClient *Client

// Redis 在中间件中初始化redis链接
func Redis() {
	db, _ := strconv.ParseUint(os.Getenv("REDIS_DB"), 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr:       os.Getenv("REDIS_ADDR"),
		Password:   os.Getenv("REDIS_PW"),
		DB:         int(db),
		MaxRetries: 1,
	})

	ctx, cancel := RedisClient.Context()
	defer cancel()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		util.Log().Panic("连接Redis不成功", err)
	}

	RedisClient = (*Client)(client)
}
