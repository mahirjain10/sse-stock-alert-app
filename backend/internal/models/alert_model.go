package models

import (
	"database/sql"
	"fmt"

	"github.com/mahirjain_10/stock-alert-app/backend/internal/types"
)

// FindAlertNameByUserIDAndAlertName retrieves a stock alert by user ID and alert name.
func FindAlertNameByUserIDAndAlertName(app *types.App, userID, alertName string) (types.StockAlert, error) {
	var stockAlert types.StockAlert

	stmt := `
	SELECT 
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
	WHERE user_id = ? AND alert_name = ?;
	`

	err := app.DB.QueryRow(stmt, userID, alertName).Scan(
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
			return types.StockAlert{}, nil // Return empty if no rows found
		}
		fmt.Printf("Error while fetching data: %v\n", err)
		return types.StockAlert{}, err
	}

	return stockAlert, nil
}

// FindAlertNameByUserIDAndID retrieves a stock alert by user ID and alert ID.
func FindAlertNameByUserIDAndID(app *types.App, userID, alertID string) (types.StockAlert, error) {
	var stockAlert types.StockAlert

	stmt := `
	SELECT 
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

	err := app.DB.QueryRow(stmt, userID, alertID).Scan(
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
			return types.StockAlert{}, nil // Return empty if no rows found
		}
		fmt.Printf("Error while fetching data: %v\n", err)
		return types.StockAlert{}, err
	}

	return stockAlert, nil
}

// InsertStockAlertData inserts a new stock alert record into the database.
func InsertStockAlertData(app *types.App, stockAlertData types.StockAlert) error {
	query := `
	INSERT INTO stock_alert (
		user_id, id, alert_name, ticker, 
		current_fetched_price, current_fetched_time, 
		alert_condition, alert_price
	)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`

	stmt, err := app.DB.Prepare(query)
	if err != nil {
		fmt.Printf("Error preparing statement: %v\n", err)
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

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
		fmt.Printf("Error executing insert: %v\n", err)
		return fmt.Errorf("failed to insert stock alert data: %w", err)
	}

	return nil
}

// UpdateStockAlertData updates an existing stock alert record in the database.
func UpdateStockAlertData(app *types.App, updateData types.UpdateStockAlert) error {
	query := `
	UPDATE stock_alert
	SET alert_name = ?, alert_condition = ?, alert_price = ?
	WHERE user_id = ? AND id = ?
	`

	stmt, err := app.DB.Prepare(query)
	if err != nil {
		fmt.Printf("Error preparing statement: %v\n", err)
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		updateData.AlertName,
		updateData.Condition,
		updateData.AlertPrice,
		updateData.UserID,
		updateData.ID,
	)

	if err != nil {
		fmt.Printf("Error executing update: %v\n", err)
		return fmt.Errorf("failed to update stock alert data: %w", err)
	}

	return nil
}

// DeleteStockAlertByID deletes a stock alert by its ID.
func DeleteStockAlertByID(app *types.App, alertID string) (int64, error) {
	tx, err := app.DB.Begin()
	if err != nil {
		return 0, fmt.Errorf("failed to begin transaction: %w", err)
	}

	query := `DELETE FROM stock_alert WHERE id = ?`
	result, err := tx.Exec(query, alertID)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("failed to delete stock alert data: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("failed to retrieve affected rows: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return 0, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return rowsAffected, nil
}

// UpdateActiveStatusByID updates the 'is_active' status of a stock alert by ID.
func UpdateActiveStatusByID(app *types.App, status bool, alertID string) error {
	query := `
	UPDATE stock_alert
	SET is_active = ?
	WHERE id = ?
	`

	stmt, err := app.DB.Prepare(query)
	if err != nil {
		fmt.Printf("Error preparing statement: %v\n", err)
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(status, alertID)
	if err != nil {
		fmt.Printf("Error executing update: %v\n", err)
		return fmt.Errorf("failed to update stock alert status: %w", err)
	}

	return nil
}

func InsertMonitorStockData(app *types.App, MSP types.MonitorStockPrice) error {
	query:=`
		INSERT INTO monitor_stock(
			id, alert_id, ticker, is_active
		)
		VALUES (?, ?, ?, ?) 
	`
	stmt, err := app.DB.Prepare(query)
	if err != nil {
		fmt.Printf("Error preparing statement: %v\n", err)
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()
	_,err = stmt.Exec(
		MSP.ID,
		MSP.AlertID,
		MSP.TickerToMonitor,
		MSP.IsActive,
	)
	if err != nil {
		fmt.Printf("Error executing insert: %v\n", err)
		return fmt.Errorf("failed to insert monitor stock data: %w", err)
	}
	return nil
}