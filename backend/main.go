package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mahirjain_10/stock-alert-app/backend/internal/app"
	"github.com/mahirjain_10/stock-alert-app/backend/internal/types"
	"github.com/mahirjain_10/stock-alert-app/backend/internal/utils"
	"github.com/mahirjain_10/stock-alert-app/backend/internal/websocket"
	"github.com/mahirjain_10/stock-alert-app/backend/web/cmd/router"
)

func main() {
	// Initialize Gin router
	r := gin.Default()
	err := app.InitalizeEnv()
	if err != nil{
		log.Fatalf("Error loading .env file: %s", err)
		return
	}
	ctx := context.Background()
	// Initialize the database and Redis client using the new helper function
	db, redisClient, err := app.InitializeServices()
	if err != nil {
		log.Fatalf("Error initializing services: %v", err)
		return
	}
	defer db.Close()
	var appInstance = types.App{
		DB:          db,
		RedisClient: redisClient,
	}
	// Initialize database tables
	if err := app.InitializeDatabaseTables(db); err != nil {
		log.Fatalf("Error initializing database tables: %v", err)
		return
	}

	hub := websocket.NewHub()

	go hub.Run()
	// Register routes
	go func() {
		log.Println("Starting Redis subscription...")
		utils.Subscribe(appInstance.RedisClient, ctx)
	}()
	router.RegisterRoutes(r, hub, &appInstance)
	log.Fatal(r.Run(":8080"))
}
