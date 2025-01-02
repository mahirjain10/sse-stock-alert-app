package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)


func InitDB() (*sql.DB, error) {
	// Connect to the database
	db, err := sql.Open("mysql", os.Getenv("SQL_DB_URL"))
	if err != nil {
		return nil, fmt.Errorf("error while connecting to database: %v", err)
	}

	// Ping the database to check if the connection is successful
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error while pinging the database: %v", err)
	}

	// Successfully connected to the database
	fmt.Println("connected to DB successfully")
	return db, nil // Return the DB connection without closing it
}
