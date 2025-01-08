package alert

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mahirjain_10/stock-alert-app/backend/internal/helpers"
	"github.com/mahirjain_10/stock-alert-app/backend/internal/models"
	"github.com/mahirjain_10/stock-alert-app/backend/internal/types"
	"github.com/mahirjain_10/stock-alert-app/backend/internal/utils"
)

func GetCurrentStockPriceAndTime(c *gin.Context, r *gin.Engine, app *types.App) {
	// var stock types.GetCurrentPrice
	var TTM types.Ticker
	var stockData types.StockData

	if !helpers.BindAndValidateJSON(c, &TTM) {
		return
	}
	latestPrice, currentTime, err := utils.GetCurrentStockPriceAndTime(TTM, stockData)
	if err != nil {
		if err.Error() == "failed to fetch stock price, try again" {
			helpers.SendResponse(c, http.StatusInternalServerError, "Failed to fetch stock price,try again", nil, nil, false)
		}
		if err.Error() == "failed to decode response json" {
			c.JSON(http.StatusBadRequest, gin.H{"statusCode": http.StatusBadRequest, "message": "failed to decode response json", "error": err.Error()})
		}
		if err.Error() == "no data found" {
			c.JSON(http.StatusNotFound, gin.H{"statusCode": http.StatusNotFound, "message": "no data found", "error": nil})
		}
	}
	// Prepare response
	response := map[string]any{
		"statusCode": http.StatusOK,
		"message":    "Latest price fetched successfully",
		"data": types.GetCurrentPrice{
			CurrentFetchedPrice: latestPrice,
			CurrentFetchedTime:  currentTime,
		},
		"error": nil,
	}
	
	// Return the response
	fmt.Println(stockData)
	// c.JSON(http.StatusOK, response)
	helpers.SendResponse(c, http.StatusOK, "Current price fetched successfully", response, nil, true)
}

func CreateStockAlert(c *gin.Context, r *gin.Engine, app *types.App) {
	ctx := context.Background()
	var alertInput types.StockAlert
	var monitorStockPrice types.MonitorStockPrice

	// Bind and validate JSON input
	if !helpers.BindAndValidateJSON(c, &alertInput) {
		return
	}

	// Check if the user exists
	user, err := models.FindUserByID(app, alertInput.UserID)
	if err != nil {
		log.Printf("Error finding user by ID: %v", err)
		helpers.SendResponse(c, http.StatusInternalServerError, "Internal server error", nil, nil, false)
		return
	}

	if user.ID == "" {
		helpers.SendResponse(c, http.StatusNotFound, "User not found", nil, nil, false)
		return
	}

	// Check if alert name already exists for the user
	existingAlert, err := models.FindAlertNameByUserIDAndAlertName(app, alertInput.UserID, alertInput.AlertName)
	if err != nil {
		log.Printf("Error finding alert name: %v", err)
		helpers.SendResponse(c, http.StatusInternalServerError, "Internal server error", nil, nil, false)
		return
	}

	if existingAlert.ID != "" {
		helpers.SendResponse(c, http.StatusConflict, "Alert name already exists. Use a different name.", nil, nil, false)
		return
	}

	if alertInput.CurrentFetchedPrice == alertInput.AlertPrice {
		helpers.SendResponse(c, http.StatusBadRequest, "Alert price cannot be same as current price", nil, nil, false)
		return
	}
	// Generate a unique ID for the alert
	alertInput.ID = uuid.New().String()

	// Insert stock alert data into the database
	if err := models.InsertStockAlertData(app, alertInput); err != nil {
		log.Printf("Error inserting stock alert data: %v", err)
		helpers.SendResponse(c, http.StatusInternalServerError, "Error saving stock alert", nil, nil, false)
		return
	}
		// Save alert data in Redis
	alertData := map[string]interface{}{
		"user_id":         user.ID,
		"ticker":          alertInput.TickerToMonitor,
		"alert_price":     alertInput.AlertPrice,
		"alert_condition": alertInput.Condition,
		"active":          alertInput.Active,
	}
	val, err := app.RedisClient.HSet(ctx, alertInput.ID, alertData).Result()
	if val == 0 {
		log.Println("Data could not saved in redis")
	}
	if err != nil {
		log.Printf("Error saving alert to Redis: %v\n", err)
	}
	
	// Insert stock monitoring data into database
	monitorStockPrice.ID=uuid.NewString()
	monitorStockPrice.AlertID=alertInput.ID
	monitorStockPrice.TickerToMonitor=alertInput.TickerToMonitor
	monitorStockPrice.IsActive=true

	err = models.InsertMonitorStockData(app,monitorStockPrice)
	if err != nil {
		log.Printf("Error inserting stock monitoring data: %v", err)
		helpers.SendResponse(c, http.StatusInternalServerError, "Error saving stock monitoring data", nil, nil, false)
		return
	}
	monitorStockHashKey := "monitor_stock:" + monitorStockPrice.ID
	monitorStockRedis := make(map[string]string)
	monitorStockRedis["id"]=monitorStockPrice.ID
	monitorStockRedis["alert_id"]=monitorStockPrice.AlertID
	monitorStockRedis["ticker"]=monitorStockPrice.TickerToMonitor
	monitorStockRedis["is_active"]=strconv.FormatBool(monitorStockPrice.IsActive)



	val, err = app.RedisClient.HSet(ctx, monitorStockHashKey, monitorStockRedis).Result()
	if val == 0 {
		log.Println("Data could not saved in redis")
	}
	if err != nil {
		log.Printf("Error saving stock monitoring data to Redis: %v\n", err)
		return;
	}

	// Publish alert to Redis channel
	utils.Publish(app.RedisClient, ctx, alertInput.TickerToMonitor, alertInput.ID)

	// Send success response
	helpers.SendResponse(c, http.StatusCreated, "Stock alert created successfully", nil, nil, true)
}

func UpdateStockAlert(c *gin.Context, r *gin.Engine, app *types.App) {
	ctx := context.Background()
	var updateAlertInput types.UpdateStockAlert
	if !helpers.BindAndValidateJSON(c, &updateAlertInput) {
		return
	}

	// Check if user exists or not
	user, err := models.FindUserByID(app, updateAlertInput.UserID)
	if err != nil {
		helpers.SendResponse(c, http.StatusInternalServerError, "Internal server error", nil, nil, false)
		return
	}

	if user.ID == "" {
		helpers.SendResponse(c, http.StatusNotFound, "User not found", nil, nil, false)
		return
	}

	// Checking for alert data with given ID exists
	retrieveStockAlertData, err := models.FindAlertNameByUserIDAndAlertName(app, updateAlertInput.UserID, updateAlertInput.AlertName)
	if err != nil {
		helpers.SendResponse(c, http.StatusInternalServerError, "Internal server error", nil, nil, false)
		return
	}

	fmt.Println(retrieveStockAlertData)

	//If alert name is already present in your account other than current alertID then send error
	if strings.TrimSpace(retrieveStockAlertData.ID) != strings.TrimSpace(updateAlertInput.ID) &&
		strings.TrimSpace(retrieveStockAlertData.UserID) == strings.TrimSpace(updateAlertInput.UserID) {
		fmt.Println("in if func")
		helpers.SendResponse(c, http.StatusNotFound, "Alert name already exists in your account,Use different alert name", nil, nil, false)
		return
	}

	err = models.UpdateStockAlertData(app, updateAlertInput)
	if err != nil {
		helpers.SendResponse(c, http.StatusInternalServerError, "Unable to update alert data ,Try again later", nil, nil, false)
		return
	}

	if retrieveStockAlertData.AlertPrice != updateAlertInput.AlertPrice {
		fmt.Println("in if for alert price update")
		// Update the data to redis
		val, err := app.RedisClient.HSet(ctx, updateAlertInput.ID, "alert_price", updateAlertInput.AlertPrice).Result()
		if val == 0 {
			log.Println("Data could not saved in redis")
		}
		if err != nil {
			// Log the error and return it or handle it as per your application's error handling policy
			log.Printf("Error updating alert in Redis for ID %s: %v", updateAlertInput.ID, err)
		}
	}
	if retrieveStockAlertData.Condition != updateAlertInput.Condition {
		fmt.Println("in if for alert condition update")
		// Update the data to redis
		val, err := app.RedisClient.HSet(ctx, updateAlertInput.ID, "alert_condition", updateAlertInput.Condition).Result()
		if val == 0 {
			log.Println("Data could not saved in redis")
		}
		if err != nil {
			// Log the error and return it or handle it as per your application's error handling policy
			log.Printf("Error updating alert in Redis for ID %s: %v", updateAlertInput.ID, err)
		}

	}
	helpers.SendResponse(c, http.StatusOK, "Stock alert updated successfully", nil, nil, true)
}

func DeleteStockAlert(c *gin.Context, r *gin.Engine, app *types.App) {
	ctx := context.Background()

	var deleteStockAlert types.DeleteStockAlert
	if !helpers.BindAndValidateJSON(c, &deleteStockAlert) {
		return
	}

	user, err := models.FindUserByID(app, deleteStockAlert.UserID)
	if err != nil {
		helpers.SendResponse(c, http.StatusInternalServerError, "Internal server error", nil, nil, false)
		return
	}

	if user.ID == "" {
		helpers.SendResponse(c, http.StatusNotFound, "User not found", nil, nil, false)
		return
	}
	retrieveStockAlertData, err := models.FindAlertNameByUserIDAndID(app, deleteStockAlert.UserID, deleteStockAlert.ID)
	if err != nil {
		helpers.SendResponse(c, http.StatusInternalServerError, "Internal server error", nil, nil, false)
		return
	}

	fmt.Println(retrieveStockAlertData)
	if retrieveStockAlertData.ID == "" {
		helpers.SendResponse(c, http.StatusNotFound, "Alert with given ID not found", nil, nil, false)
		return
	}

	rowsAffected, err := models.DeleteStockAlertByID(app, retrieveStockAlertData.UserID)
	if err != nil {
		helpers.SendResponse(c, http.StatusInternalServerError, "Internal server error", nil, nil, false)
		return
	}
	if rowsAffected == 0 {
		helpers.SendResponse(c, http.StatusNotFound, "Stock Alert to be deleted not found", nil, nil, false)
		return
	} else {
		_, err := app.RedisClient.Del(ctx, retrieveStockAlertData.ID).Result()
		if err != nil {
			log.Printf("Error deleting alert in Redis for ID %s: %v", retrieveStockAlertData.ID, err)
		}
		helpers.SendResponse(c, http.StatusOK, "Stock Alert deleted successfully", nil, nil, true)
		return
	}
}

func UpdateActiveStatus(c *gin.Context, r *gin.Engine, app *types.App) {
	ctx := context.Background()

	var updateActiveStatus types.UpdateActiveStatus
	if !helpers.BindAndValidateJSON(c, &updateActiveStatus) {
		return
	}

	isSuccess := utils.UpdateActiveStatusUtil(c,ctx,updateActiveStatus.UserID,updateActiveStatus.ID,updateActiveStatus.Active,app)

	if isSuccess {
		helpers.SendResponse(c, http.StatusOK, "Stock alert status updated successfully", nil, nil, true)
	}
}
