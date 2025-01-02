package types

import (
	"database/sql"
	"time"

	"github.com/redis/go-redis/v9"
)

type App struct {
	DB          *sql.DB
	RedisClient *redis.Client
}

type RegisterUser struct {
	ID        string    `json:"id,omitempty"` 
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"` 
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type EditUser struct {
	ID          string `json:"id" binding:"required"`
	Name        string `json:"name,omitempty"`
	Email       string `json:"email,omitempty"`
	OldPassword string `json:"old_password" binding:"required_with=NewPassword"`
	NewPassword string `json:"new_password" binding:"required_with=OldPassword"`
}

type LoginUser struct {
	ID       string `json:"id,omitempty"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Ticker struct {
	TickerToMonitor string `json:"ticker_to_monitor"`
}
type GetCurrentPrice struct {
	CurrentFetchedPrice float64 `json:"current_fetched_price"`
	CurrentFetchedTime  string  `json:"current_fetched_time"`
	AlertID             string
}

type MonitorStockPrice struct {
	Ticker
	AlertID string `json:"alert_id"`
}
type StockAlert struct {
	UserID    string `json:"user_id"`
	ID        string `json:"id,omitempty"`
	AlertName string `json:"alert_name"`
	Ticker
	GetCurrentPrice
	Condition  string  `json:"alert_condition"`
	AlertPrice float64 `json:"alert_price"`
	Active     bool    `json:"active"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type UpdateStockAlert struct {
	UserID     string  `json:"user_id"`
	ID         string  `json:"id,omitempty"`
	AlertName  string  `json:"alert_name"`
	Condition  string  `json:"alert_condition"`
	AlertPrice float64 `json:"alert_price"`
}

type UpdateActiveStatus struct {
	UserID string `json:"user_id"`
	ID     string `json:"id"`
	Active bool   `json:"active"`
}

type DeleteStockAlert struct {
	UserID string `json:"user_id"`
	ID     string `json:"id"`
}
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
