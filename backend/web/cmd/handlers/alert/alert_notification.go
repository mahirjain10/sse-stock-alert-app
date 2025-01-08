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
	if !helpers.BindAndValidateJSON(c,UpdateActiveStatus) {
		return
	}
	
	utils.UpdateActiveStatusUtil(c,ctx,UpdateActiveStatus.UserID,UpdateActiveStatus.ID,UpdateActiveStatus.Active,app)
	
	data  := make(map[string]interface{})

    // Add entries to the map
    data["user_id"] = UpdateActiveStatus.UserID
    data["alert_id"] = UpdateActiveStatus.ID
	helpers.SendResponse(c,http.StatusOK,"Alert condition met",data,nil,true)
}


