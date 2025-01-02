package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

// InitializeRedis sets up the Redis client
func InitializeRedis() (*redis.Client,error) {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_DB_URL"),
		Password: os.Getenv("REDIS_DB_PASSWORD"), 
		DB:       0,       
	})

	// Ping Redis to test connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
		return nil,fmt.Errorf("Failed to connect to Redis: %v",err)
	}

	fmt.Println("Redis connection established")
	return client,nil
}
