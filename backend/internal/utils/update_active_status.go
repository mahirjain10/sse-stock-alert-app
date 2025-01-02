package utils

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mahirjain_10/stock-alert-app/backend/internal/helpers"
	"github.com/mahirjain_10/stock-alert-app/backend/internal/models"
	"github.com/mahirjain_10/stock-alert-app/backend/internal/types"
	// model "github.com/mahirjain_10/stock-alert-app/backend/internal/models"
)

func UpdateActiveStatusUtil(c *gin.Context,ctx context.Context,userID string,alertID string ,updatedStatus bool,app *types.App) {
	user, err := models.FindUserByID(app, userID)
	if err != nil {
		helpers.SendResponse(c, http.StatusInternalServerError, "Internal server error", nil, nil, false)
		return
	}
	
	if user.ID == "" {
		helpers.SendResponse(c, http.StatusNotFound, "User not found", nil, nil, false)
		return
	}

	retrieveStockAlertData, err := models.FindAlertNameByUserIDAndID(app, userID, alertID)
	if err != nil {
		helpers.SendResponse(c, http.StatusInternalServerError, "Internal sever error", nil, nil, false)
		return
	}

	fmt.Println(retrieveStockAlertData)
	if retrieveStockAlertData.ID == "" {
		helpers.SendResponse(c, http.StatusNotFound, "Alert with given ID not found", nil, nil, false)
		return
	}

	err = models.UpdateActiveStatusByID(app, updatedStatus, alertID)
	if err != nil {
		helpers.SendResponse(c, http.StatusInternalServerError, "Unable to update alert status ,Try again later", nil, nil, false)
		return
	}

	if retrieveStockAlertData.Active != updatedStatus {
		val, err := app.RedisClient.HSet(ctx, userID, "active", updatedStatus).Result()
		if val == 0 {
			log.Println("Could not save alert status in redis")
		}
		if err != nil {
			// Log the error and return it or handle it as per your application's error handling policy
			log.Printf("Error updating alert status in Redis for ID %s: %v", userID, err)
		}
	}
}