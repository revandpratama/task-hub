package config

import "github.com/redis/go-redis/v9"

var RedisClient *redis.Client

func InitRedisClient() {
	rds := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	RedisClient = rds
}
