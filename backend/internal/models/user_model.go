package models

import (
	"database/sql"
	"fmt"

	"github.com/mahirjain_10/stock-alert-app/backend/internal/types"
)

func FindUserByID(app *types.App, user_id string)(types.RegisterUser, error){
	stmt := `SELECT * FROM user WHERE id = ?`

	var user types.RegisterUser

	err := app.DB.QueryRow(stmt, user_id).Scan(&user.ID, &user.Name, &user.Email, &user.Password,&user.CreatedAt,&user.UpdatedAt) // assuming the columns you are selecting
	if err != nil {
        if err == sql.ErrNoRows {
            // No user found with the given email
            return types.RegisterUser{}, nil
        }
        // Other errors
        fmt.Printf("Error while fetching data: %v\n", err)
        return types.RegisterUser{}, err
    }

    return user, nil
}

func FindUserByEmail(app *types.App, email string) (types.RegisterUser, error) {
    stmt := `SELECT * FROM user WHERE email = ?`

    var user types.RegisterUser

    // Use QueryRow() to execute the SELECT statement and get a single row
    err := app.DB.QueryRow(stmt, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password,&user.CreatedAt,&user.UpdatedAt) // assuming the columns you are selecting
    if err != nil {
        if err == sql.ErrNoRows {
            // No user found with the given email
            return types.RegisterUser{}, nil
        }
        // Other errors
        fmt.Printf("Error while fetching data: %v\n", err)
        return types.RegisterUser{}, err
    }

    return user, nil
}

func InsertUser(app *types.App, user types.RegisterUser) error {
	// SQL query to insert a new user
	query := `
		INSERT INTO user(id, name, email, password) 
		VALUES(?, ?, ?, ?)
	`

	// Prepare the statement
	stmt, err := app.DB.Prepare(query)
	if err != nil {
		// Log and return error if statement preparation fails
		fmt.Printf("error preparing statement: %v\n", err)
		return fmt.Errorf("failed to prepare statement: %w", err)
	}

	// Ensure the statement is closed after execution
	defer stmt.Close()

	// Execute the prepared statement with user data
	_, err = stmt.Exec(user.ID, user.Name, user.Email, user.Password)
	if err != nil {
		// Log and return error if execution fails
		fmt.Printf("error while inserting user data: %v\n", err)
		return fmt.Errorf("failed to insert user: %w", err)
	}

	// Return nil if the operation was successful
	return nil
}

