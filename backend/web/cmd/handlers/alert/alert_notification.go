package alert

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mahirjain_10/stock-alert-app/backend/internal/helpers"
	"github.com/mahirjain_10/stock-alert-app/backend/internal/types"
	"github.com/mahirjain_10/stock-alert-app/backend/internal/utils"
)

func SendAlertNotification(c *gin.Context, r *gin.Engine, app *types.App) {
	ctx := context.Background()
	
	var UpdateActiveStatus types.UpdateActiveStatus
	if !helpers.BindAndValidateJSON(c, &UpdateActiveStatus) {
		return
	}
	
	// Unregister the WebSocket client before updating the status
	// if hub := app.Hub; hub != nil {
	// 	hub.UnregisterClientByAlertID(UpdateActiveStatus.ID)
	// }
	
	utils.UpdateActiveStatusUtil(c, ctx, UpdateActiveStatus.UserID, UpdateActiveStatus.ID, UpdateActiveStatus.Active, app)
	
	data := map[string]interface{}{
		"user_id": UpdateActiveStatus.UserID,
		"alert_id": UpdateActiveStatus.ID,
	}
	
	helpers.SendResponse(c, http.StatusOK, "Alert condition met", data, nil, true)
}


