package cacheRedis

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type RedisClientInterface interface {
	Set(ctx context.Context, key, value string)
	Get(ctx context.Context, key string) string
}

type redisClient struct {
	addr     string
	password string
	db       int
}

func NewRedisClient() RedisClientInterface {
	return &redisClient{
		addr:     "localhost:6379",
		password: "",
		db:       0,
	}
}

func (rc *redisClient) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     rc.addr,
		Password: rc.password,
		DB:       rc.db,
	})
}

func (rc *redisClient) Get(ctx context.Context, key string) string {
	client := rc.getClient()
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		val = ""
	}

	return val
}

func (rc *redisClient) Set(ctx context.Context, key, value string) {
	client := rc.getClient()

	err := client.Set(ctx, key, value, 0).Err()
	if err != nil {
		panic(err)
	}
}
