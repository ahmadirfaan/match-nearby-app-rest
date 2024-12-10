package storage

import (
	"context"
	"github.com/ahmadirfaan/match-nearby-app-rest/app"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

func InitRedis() *redis.Client {

	appConfig := app.Init().Config
	redisClient := redis.NewClient(&redis.Options{
		Addr: appConfig.RedisAddress,
	})

	ctx := context.Background()

	if err := redisClient.Ping(ctx).Err(); err != nil {
		logrus.Fatalf("Failed to connect to Redis: %v", err)
	}
	logrus.Println("Connected to Redis!")
	return redisClient
}
