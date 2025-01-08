package app

import (
	"database/sql"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/mahirjain_10/stock-alert-app/backend/internal/database"
	"github.com/mahirjain_10/stock-alert-app/backend/internal/models"
	"github.com/redis/go-redis/v9"
)

// InitializeServices initializes the database and Redis client
func InitializeServices() (*sql.DB,*redis.Client,error) {
	// Initialize the database connection
	db, err := database.InitDB()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to initialize database: %w", err)
	}

	// Initialize the Redis client
	redisClient, err := database.InitializeRedis()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to initialize Redis: %w", err)
	}

	return db, redisClient, nil
}

// InitializeDatabaseTables initializes required database tables (user and stock alerts)
func InitializeDatabaseTables(db *sql.DB) error {
	// Initialize user table
	if err := models.InitUserTable(db); err != nil {
		return fmt.Errorf("error creating user table: %w", err)
	}

	// Initialize stock alert table
	if err := models.InitStockAlertTable(db); err != nil {
		return fmt.Errorf("error creating stock alert table: %w", err)
	}

	// Initalize monitor stock table
	if err := models.InitializeMonitorStockTable(db); err != nil{
		return fmt.Errorf("error creating monitor stock table: %w",err)
	}

	return nil
}


// InitializeEnv initializes envs and returns error 
func InitalizeEnv() error {
	err := godotenv.Load(".env")
	if err != nil{
		return fmt.Errorf("error initalizing env: %w",err)
	}
	return nil
}


