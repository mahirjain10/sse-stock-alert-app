package utils

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mahirjain_10/stock-alert-app/backend/internal/helpers"
	"github.com/mahirjain_10/stock-alert-app/backend/internal/models"
	"github.com/mahirjain_10/stock-alert-app/backend/internal/types"
	// model "github.com/mahirjain_10/stock-alert-app/backend/internal/models"
)

func UpdateActiveStatusUtil(c *gin.Context, ctx context.Context, userID string, alertID string, updatedStatus bool, app *types.App) bool {

	// Check if user exists or not
	user, err := models.FindUserByID(app, userID)
	if err != nil {
		fmt.Printf("error in updateStatusUtils : %v",err)
		helpers.SendResponse(c, http.StatusInternalServerError, "Internal server error", nil, nil, false)
		return false
	}

	if user.ID == "" {
		fmt.Printf("error in updateStatusUtils : %v",err)
		helpers.SendResponse(c, http.StatusNotFound, "User not found", nil, nil, false)
		return false
	}

	// Check if stock data ID mapped to userID
	retrieveStockAlertData, err := models.FindAlertNameByUserIDAndID(app, userID, alertID)
	if err != nil {
		fmt.Printf("error in updateStatusUtils : %v",err)
		helpers.SendResponse(c, http.StatusInternalServerError, "Internal sever error", nil, nil, false)
		return false
	}

	fmt.Println(retrieveStockAlertData)
	if retrieveStockAlertData.ID == "" {
		helpers.SendResponse(c, http.StatusNotFound, "Alert with given ID not found", nil, nil, false)
		return false
	}

	// update alert status in stock alert db
	err = models.UpdateActiveStatusByID(app, updatedStatus, alertID)
	if err != nil {
		fmt.Printf("error in updateStatusUtils : %v",err)
		helpers.SendResponse(c, http.StatusInternalServerError, "Unable to update alert status ,Try again later", nil, nil, false)
		return false
	}
     
	// IF alert status in DB is not equal to Update Alert Status then
	if retrieveStockAlertData.Active != updatedStatus {
		fmt.Printf("updated Status : %t and alertID : %s", updatedStatus, alertID)
		statusStr := strconv.FormatBool(updatedStatus)
		_, err := app.RedisClient.HSet(ctx, alertID, "active", statusStr).Result()
		if err != nil {
			log.Printf("Error updating alert status in Redis for ID %s: %v", alertID, err)
			return false
		}
		
		// Verify if the update was successful
		storedValue, err := app.RedisClient.HGet(ctx, alertID, "active").Result()
		if err != nil {
			log.Printf("Error retrieving alert status from Redis for ID %s: %v", alertID, err)
			return false
		}
		
		// Check if the retrieved value matches the expected value
		if storedValue != statusStr {
			log.Printf("Redis update verification failed: expected %s, got %s", statusStr, storedValue)
			return false
		}
		
		log.Println("Redis alert status updated successfully")
		return true

	}
	return true
}
