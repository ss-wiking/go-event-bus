package app

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
	"strconv"
)

func InitRedisClient() *redis.Client {
	address := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	db, _ := strconv.Atoi(os.Getenv("REDIS_DB"))

	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: os.Getenv("REDIS_PASS"),
		DB:       db,
	})

	err := client.Ping(context.Background()).Err()

	if err != nil {
		panic(err)
	}

	return client
}
