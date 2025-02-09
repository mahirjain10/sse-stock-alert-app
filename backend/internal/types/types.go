package types

import (
	"database/sql"
	"time"

	// "github.com/mahirjain_10/stock-alert-app/backend/internal/websocket"

	// "github.com/mahirjain_10/stock-alert-app/backend/internal/websocket"

	// "github.com/mahirjain_10/stock-alert-app/backend/internal/websocket"
	"github.com/golang-jwt/jwt"
	"github.com/redis/go-redis/v9"
)

// App contains the database and Redis client
type App struct {
	DB          *sql.DB
	RedisClient *redis.Client
	// Hub *websocket.Hub
}

// -----------------------------------------
// User Management Types
// -----------------------------------------

// RegisterUser represents a user during registration
type RegisterUser struct {
	ID        string    `json:"id,omitempty"` 
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"` 
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// EditUser represents a user's data for editing their profile
type EditUser struct {
	ID          string `json:"id" binding:"required"`
	Name        string `json:"name,omitempty"`
	Email       string `json:"email,omitempty"`
	OldPassword string `json:"old_password" binding:"required_with=NewPassword"`
	NewPassword string `json:"new_password" binding:"required_with=OldPassword"`
}

// LoginUser represents the data required for user login
type LoginUser struct {
	ID       string `json:"id,omitempty"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// -----------------------------------------
// Stock Monitoring Types
// -----------------------------------------

// Ticker represents a stock ticker to monitor
type Ticker struct {
	TickerToMonitor string `json:"ticker_to_monitor"`
}

// GetCurrentPrice represents the current price of a stock along with its fetch time
type GetCurrentPrice struct {
	CurrentFetchedPrice float64 `json:"current_fetched_price"`
	CurrentFetchedTime  string  `json:"current_fetched_time"`
	AlertID             string  `json:"alert_id"`
}

// MonitorStockPrice represents stock monitoring data with an alert ID
type MonitorStockPrice struct {
	ID string 
	Ticker
	AlertID string `json:"alert_id"`
	IsActive bool `json:"is_active"`
}

// -----------------------------------------
// Stock Alert Types
// -----------------------------------------

// StockAlert represents a user's stock alert data
type StockAlert struct {
	UserID    string    `json:"user_id"`
	ID        string    `json:"id,omitempty"`
	AlertName string    `json:"alert_name"`
	Ticker
	GetCurrentPrice
	Condition  string  `json:"alert_condition"`
	AlertPrice float64 `json:"alert_price"`
	Active     bool    `json:"active"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type StartMonitoring struct{
	AlertID string `json:"alert_id"`
	UserID string `json:"user_id"`
	Ticker
}
// UpdateStockAlert represents the fields to update an existing stock alert
type UpdateStockAlert struct {
	UserID     string  `json:"user_id"`
	ID         string  `json:"id,omitempty"`
	AlertName  string  `json:"alert_name"`
	Condition  string  `json:"alert_condition"`
	AlertPrice float64 `json:"alert_price"`
}

// UpdateActiveStatus represents the status update of a stock alert
type UpdateActiveStatus struct {
	UserID string `json:"user_id"`
	ID     string `json:"id"`
	Active bool   `json:"active"`
}

// DeleteStockAlert represents the data required to delete a stock alert
type DeleteStockAlert struct {
	UserID string `json:"user_id"`
	ID     string `json:"id"`
}

type CustomClaims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}


// -----------------------------------------
// Stock Data Structure for Chart Data
// -----------------------------------------

// StockData represents the response structure for fetched stock chart data
type StockData struct {
	Chart struct {
		Result []struct {
			Indicators struct {
				Quote []struct {
					Close []float64 `json:"close"`
				} `json:"quote"`
			} `json:"indicators"`
		} `json:"result"`
	} `json:"chart"`
}
