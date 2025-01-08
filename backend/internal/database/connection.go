package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// InitDB establishes a connection to the MySQL database and returns the connection object
func InitDB() (*sql.DB, error) {
	// Retrieve the database URL from environment variables
	dbURL := os.Getenv("SQL_DB_URL")
	if dbURL == "" {
		return nil, fmt.Errorf("SQL_DB_URL environment variable not set")
	}

	// Connect to the MySQL database
	db, err := sql.Open("mysql", dbURL)
	if err != nil {
		return nil, fmt.Errorf("error while connecting to database: %v", err)
	}

	// Ping the database to check the connection is working
	err = db.Ping()
	if err != nil {
		// Return error if ping fails
		return nil, fmt.Errorf("error while pinging the database: %v", err)
	}

	// Log successful connection
	fmt.Println("Connected to the database successfully")
	return db, nil
}
