package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mahirjain_10/stock-alert-app/backend/internal/types"
	"github.com/mahirjain_10/stock-alert-app/backend/internal/websocket"
	"github.com/mahirjain_10/stock-alert-app/backend/web/cmd/handlers/alert"
	"github.com/mahirjain_10/stock-alert-app/backend/web/cmd/handlers/auth"
)

// registerRoutes handles the grouping and organization of routes
func RegisterRoutes(r *gin.Engine, hub *websocket.Hub, app *types.App) {
	// WebSocket endpoint
	r.GET("/ws/get-stock-price-socket", func(c *gin.Context) {
		websocket.ServeWs(c, hub, c.Writer, c.Request)
	})

	// Auth group
	authRoutes := r.Group("/api/auth")
	{
		authRoutes.POST("/register", func(c *gin.Context) {
			auth.RegisterUser(c, r, app)
		})
		authRoutes.POST("/login", func(c *gin.Context) {
			auth.LoginUser(c, r, app)
		})
	}

	// Alert group
	alertRoutes := r.Group("/api/alert")
	{
		alertRoutes.POST("/get-current-price", func(c *gin.Context) {
			alert.GetCurrentStockPriceAndTime(c, r, app)
		})
		alertRoutes.POST("/create-stock-alert", func(c *gin.Context) {
			alert.CreateStockAlert(c, r, app)
		})
		alertRoutes.PUT("/update-stock-alert", func(c *gin.Context) {
			alert.UpdateStockAlert(c, r, app)
		})
		alertRoutes.PUT("/update-stock-alert-status", func(c *gin.Context) {
			alert.UpdateActiveStatus(c, r, app)
		})
		alertRoutes.DELETE("/delete-stock-alert", func(c *gin.Context) {
			alert.DeleteStockAlert(c, r, app)
		})
	}
}