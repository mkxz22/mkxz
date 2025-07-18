package redis

import "github.com/go-redis/redis/v8"

var Client *redis.Client

func Redis(addr string, password string, db int) {
	Client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
}
