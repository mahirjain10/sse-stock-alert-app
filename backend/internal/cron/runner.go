package cron

import (
	"time"

	"github.com/mahirjain_10/stock-alert-app/backend/internal/types"
	"github.com/mahirjain_10/stock-alert-app/backend/internal/websocket"
	"github.com/robfig/cron/v3"
)

func StartCron(app *types.App, hub *websocket.Hub) *cron.Cron {
	// Set timezone to IST (New Delhi)
	ist, _ := time.LoadLocation("Asia/Kolkata")
	c := cron.New(cron.WithLocation(ist))

	// Start monitoring at 2:27 PM IST
	c.AddFunc("52 19 * * *", func() {
		StartMonitoringWithRetry(app)
	})

	// Stop monitoring at 2:27:30 PM IST
	c.AddFunc("48 14 * *", func() {
		go func() {
			time.Sleep(30 * time.Second) // Wait for 30 seconds before stopping
			StopMonitoringJob(app, hub)
		}()
	})

	c.Start()
	return c
}
