package app

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/mahirjain_10/stock-alert-app/backend/internal/database"
	"github.com/mahirjain_10/stock-alert-app/backend/internal/helpers"
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


func InitializeLogger() (*os.File,error) {
	// Open or create a log file
	file, err := os.OpenFile("C:\\Users\\Mahir\\Documents\\stock app\\backend\\internal\\logs\\app.log", os.O_CREATE|os.O_WRONLY | os.O_TRUNC, 0666)
	if err != nil {
		slog.Error("Failed to open log file", "error", err)
		return nil ,err
	}

	// Create a JSON logger
	logger := slog.New(slog.NewJSONHandler(file, nil))

	// Attach file and line number
	childLogger := logger.With(slog.Group("file info", "file", helpers.GetFileName(), "line", helpers.GetLineNumber()))

	// Set as default logger
	slog.SetDefault(childLogger)
	return file ,nil 
}



