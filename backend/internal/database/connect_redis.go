package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

// InitializeRedis sets up and returns a Redis client
func InitializeRedis() (*redis.Client, error) {
	// Load Redis connection details from environment variables
	redisAddr := os.Getenv("REDIS_DB_URL")
	redisPassword := os.Getenv("REDIS_DB_PASSWORD")

	// Create a new Redis client
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword, // empty password means no password
		DB:       0,             // default DB
	})

	// Ping Redis to test the connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		// Log the error and return it
		log.Printf("Failed to connect to Redis: %v", err)
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	// Log success and return the Redis client
	log.Println("Redis connection established")
	return client, nil
}
