package models

import (
	"fmt"

	"database/sql"
)

func InitUserTable(db *sql.DB) error {
	// TODO : change "users" to "user"
	query := `
		CREATE TABLE IF NOT EXISTS user (
			id VARCHAR(36) PRIMARY KEY,
			name VARCHAR(20) NOT NULL,
			email VARCHAR(30) UNIQUE NOT NULL,
			password CHAR(60) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
		)
	`

	// Execute the query
	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	fmt.Println("User table created successfully!")
	return nil
}

func InitStockAlertTable(db *sql.DB) error{
	query := `
	CREATE TABLE IF NOT EXISTS stock_alert (
    user_id VARCHAR(36) NOT NULL,
    id VARCHAR(36) PRIMARY KEY,
    alert_name VARCHAR(50) NOT NULL,
    ticker VARCHAR(20) NOT NULL, -- Assuming ticker is a single value; adjust datatype as needed
    current_fetched_price DECIMAL(6,2) NOT NULL,
    current_fetched_time DATETIME NOT NULL, 
    alert_condition VARCHAR(2) NOT NULL, -- Assuming alert_condition is a string; adjust length if needed
    alert_price DECIMAL(6,2) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE NOT NULL,
    created_on DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_on DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE
	)	
	`
	// Execute the query
	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	fmt.Println("Stock alert table created successfully!")
	return nil
}