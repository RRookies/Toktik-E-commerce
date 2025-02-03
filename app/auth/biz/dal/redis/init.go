package redis

import (
	"Toktik-E-commerce/app/auth/conf"
	"context"
	"github.com/redis/go-redis/v9"
)

func NewRedisClient(ctx context.Context) *redis.Client {
	c := ctx.Value("config").(*conf.Config)
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     c.GetConfig().Redis.Address,
		Username: c.GetConfig().Redis.Username,
		Password: c.GetConfig().Redis.Password,
		DB:       c.GetConfig().Redis.DB,
	})
	if err := RedisClient.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
	return RedisClient
}
