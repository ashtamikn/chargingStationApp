package cache

import (
	"charging-assignment1/models"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type redisCache struct {
	host    string
	db      int
	expires time.Duration
}

// type redisCachee struct {
// 	host    string
// 	db      int
// 	expires time.Duration
// }

func NewRedisCache(host string, db int, exp time.Duration) ChargingStationCache {
	return &redisCache{
		host:    host,
		db:      db,
		expires: exp,
	}
}

//	func NewRedisCachee(host string, db int, exp time.Duration) StartChargingCache {
//		return &redisCachee{
//			host:    host,
//			db:      db,
//			expires: exp,
//		}
//	}
func (cache *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "",
		DB:       cache.db,
	})
}
func (cache *redisCache) Set(key string, value []*models.ChargingStation) {
	client := cache.getClient()
	json, err := json.Marshal(&value)
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	client.Set(ctx, key, json, cache.expires*time.Second)
}
func (cache *redisCache) Get(key string) []*models.ChargingStation {
	client := cache.getClient()
	ctx := context.Background()
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		return nil
	}
	fmt.Print("inside redis get")
	cs := []models.ChargingStation{}
	json.Unmarshal([]byte(val), &cs)
	if err != nil {
		panic(err)
	}
	csPtr := make([]*models.ChargingStation, len(cs))
	for i := range cs {
		csPtr[i] = &cs[i]
	}

	return csPtr
}
