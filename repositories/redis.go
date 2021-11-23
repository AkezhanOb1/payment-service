package repositories

import (
	"context"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

func RetrieveMerchantApiSessionRepository(sID string) (string, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("RedisAddress"),
		Password: os.Getenv("RedisPassword"), // no password set
		DB:       0,                          // use default DB
	})

	val, err := redisClient.Get(context.Background(), sID).Result()
	if err != nil {
		return "", err
	}
	return val, err
}

func SetMerchantApiSessionRepository(sID string, session string) error {

	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("RedisAddress"),
		Password: os.Getenv("RedisPassword"), // no password set
		DB:       0,                          // use default DB
	})

	err := redisClient.Set(context.Background(), sID, session, time.Hour*3).Err()
	if err != nil {
		return err
	}

	return nil
}
