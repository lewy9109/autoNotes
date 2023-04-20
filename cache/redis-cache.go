package cache

import (
	"context"
	"encoding/json"
	"github/lewy9109/autoNotes/inspection/controller"
	"time"

	"github.com/go-redis/redis/v8"
)

type redisCache struct {
	host    string
	db      int
	expires time.Duration
}

func NewRedisCache(host string, db int, exp time.Duration) InspectionCarCache {
	return &redisCache{
		host:    host,
		db:      db,
		expires: exp,
	}
}

func (cache *redisCache) Set(key string, value *controller.GetInspectionResponse) {
	client := cache.getClient()

	json, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	ctx := context.TODO()
	client.Set(ctx, key, string(json), cache.expires*time.Second)
}

func (cache *redisCache) Get(key string) *controller.GetInspectionResponse {
	client := cache.getClient()

	ctx := context.TODO()
	val, err := client.Get(ctx, key).Result()

	if err != nil {
		return nil
	}

	inspectCar := controller.GetInspectionResponse{}
	err = json.Unmarshal([]byte(val), &inspectCar)
	if err != nil {
		panic(err)
	}

	return &inspectCar
}

func (cache *redisCache) getClient() *redis.Client {
	redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "",
		DB:       cache.db,
	})
}
