package alert

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mahirjain_10/stock-alert-app/backend/internal/helpers"
	"github.com/mahirjain_10/stock-alert-app/backend/internal/types"
)

func SendAlertNotification(c *gin.Context, r *gin.Engine, app *types.App) {

	var response types.Response
	if !helpers.BindAndValidateJSON(c, &response) {
		return
	}

	responseData, err := json.Marshal(response.Data)
	if err != nil {
		log.Printf("Error marshaling response: %v", err)
	}
	var UpdateActiveStatus types.UpdateActiveStatus
	err = json.Unmarshal(responseData, &UpdateActiveStatus)
	if err != nil {
		log.Printf("Error unmarshaling response: %v", err)
	}
	
	
	helpers.SendResponse(c,http.StatusOK,"Alert condition met",response.Data,nil,true)
}
