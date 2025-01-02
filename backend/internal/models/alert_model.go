package models

import (
	"database/sql"
	"fmt"

	"github.com/mahirjain_10/stock-alert-app/backend/internal/types"
)

func FindAlertNameByUserIDAndAlertName(app *types.App, userID string, alertName string) (types.StockAlert, error) {
	var stockAlert types.StockAlert
	var err error
	// This is to check if alert_name is unique on account level
	stmt := `SELECT 
    id,
    user_id,
    ticker,
    alert_name,
    current_fetched_price,
    current_fetched_time,
    alert_condition,
    alert_price,
    is_active,
    created_on,
    updated_on
	FROM stock_alert
	WHERE user_id = ? AND alert_name = ?;`

	err = app.DB.QueryRow(stmt, userID, alertName).Scan(
		&stockAlert.ID,
		&stockAlert.UserID,
		&stockAlert.Ticker.TickerToMonitor,
		&stockAlert.AlertName,
		&stockAlert.CurrentFetchedPrice,
		&stockAlert.CurrentFetchedTime,
		&stockAlert.Condition,
		&stockAlert.AlertPrice,
		&stockAlert.Active,
		&stockAlert.CreatedAt,
		&stockAlert.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			// If no alert found with given alert name found
			return types.StockAlert{}, nil
		}
		// Other errors
		fmt.Printf("Error while fetching data: %v\n", err)
		return types.StockAlert{}, err
	}
	return stockAlert, nil
}

func FindAlertNameByUserIDAndID(app *types.App, userID string, ID string) (types.StockAlert, error) {
	var stockAlert types.StockAlert
	var err error

	// This is to check if alert_name is not present other than
	// current alert's ID in a particular user account to avoid duplication
	// on account level while updating the alert data
	stmt := `SELECT 
    id,
    user_id,
    ticker,
    alert_name,
    current_fetched_price,
    current_fetched_time,
    alert_condition,
    alert_price,
    is_active,
    created_on,
    updated_on
	FROM stock_alert
	WHERE user_id = ? AND id = ?;
	`
	err = app.DB.QueryRow(stmt, userID, ID).Scan(
		&stockAlert.ID,
		&stockAlert.UserID,
		&stockAlert.Ticker.TickerToMonitor,
		&stockAlert.AlertName,
		&stockAlert.CurrentFetchedPrice,
		&stockAlert.CurrentFetchedTime,
		&stockAlert.Condition,
		&stockAlert.AlertPrice,
		&stockAlert.Active,
		&stockAlert.CreatedAt,
		&stockAlert.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			// If no alert found with given alert name found
			return types.StockAlert{}, nil
		}
		// Other errors
		fmt.Printf("Error while fetching data: %v\n", err)
		return types.StockAlert{}, err
	}
	return stockAlert, nil
}

func InsertStockAlertData(app *types.App, stockAlertData types.StockAlert) error {
	// SQL query to insert stock alert data
	query := `
    INSERT INTO stock_alert (user_id, id, alert_name, ticker, current_fetched_price,current_fetched_time, alert_condition, alert_price)
    VALUES (?, ?, ?, ?, ?, ?, ?,?)
    `

	// Prepare the statement
	stmt, err := app.DB.Prepare(query)
	if err != nil {
		fmt.Printf("error preparing statement: %v\n", err)
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close() // Ensure the statement is closed after use

	// Execute the prepared statement with stock alert data
	_, err = stmt.Exec(
		stockAlertData.UserID,
		stockAlertData.ID,
		stockAlertData.AlertName,
		stockAlertData.Ticker.TickerToMonitor,
		stockAlertData.CurrentFetchedPrice,
		stockAlertData.CurrentFetchedTime,
		stockAlertData.Condition,
		stockAlertData.AlertPrice,
	)
	if err != nil {
		fmt.Printf("error executing stock alert insert: %v\n", err)
		return fmt.Errorf("failed to insert stock alert data: %w", err)
	}

	return nil // Return nil if operation was successful
}

func UpdateStockAlertData(app *types.App, updateData types.UpdateStockAlert) error {
	// SQL query to update stock alert data
	query := `
		UPDATE stock_alert
		SET alert_name = ?, 
			alert_condition = ?, 
			alert_price = ?
		WHERE user_id = ? AND id = ?
	`
	// Prepare the statement
	stmt, err := app.DB.Prepare(query)
	if err != nil {
		fmt.Printf("error preparing statement: %v\n", err)
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close() // Ensure the statement is closed after use
	// Execute the prepared statement with stock alert data
	_, err = stmt.Exec(
		updateData.AlertName,
		updateData.Condition,
		updateData.AlertPrice,
		updateData.UserID,
		updateData.ID,
	)
	if err != nil {
		fmt.Printf("failed to update stock alert data: %v\n", err)
		return fmt.Errorf("failed to update stock alert data: %w", err)
	}

	return nil // Return nil if operation was successful
}

func DeleteStockAlertByID(app *types.App, ID string) (int64, error) {
	// Start a new transaction
	tx, err := app.DB.Begin()
	if err != nil {
		fmt.Printf("failed to begin transaction: %v", err)
		return 0, fmt.Errorf("failed to begin transaction: %w", err)
	}

	// Prepare the DELETE query
	query := `DELETE FROM stock_alert WHERE id=?`
	result, err := tx.Exec(query, ID)
	if err != nil {
		tx.Rollback() // Roll back the transaction if an error occurs
		fmt.Printf("failed to delete stock alert data: %v", err)
		return 0, fmt.Errorf("failed to delete stock alert data: %w", err)
	}

	// Get the number of rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		fmt.Printf("failed to retrieve affected rows: %v", err)
		return 0, fmt.Errorf("failed to retrieve affected rows: %w", err)
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		fmt.Printf("failed to commit transaction: %v", err)
		return 0, fmt.Errorf("failed to commit transaction: %w", err)
	}

	fmt.Printf("Rows affected: %d\n", rowsAffected)
	return rowsAffected, nil
}

func UpdateActiveStatusByID(app *types.App, status bool, ID string) error {
	query := `
		UPDATE stock_alert
		SET is_active = ?
		WHERE id = ?
	`
	// Prepare the statement
	stmt, err := app.DB.Prepare(query)
	if err != nil {
		fmt.Printf("error preparing statement: %v\n", err)
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close() // Ensure the statement is closed after use
	// Execute the prepared statement with stock alert data
	_, err = stmt.Exec(
		status,
		ID,
	)
	if err != nil {
		fmt.Printf("failed to update stock alert status: %v\n", err)
		return fmt.Errorf("failed to update stock alert status: %w", err)
	}

	return nil // Return nil if operation was successful
}
