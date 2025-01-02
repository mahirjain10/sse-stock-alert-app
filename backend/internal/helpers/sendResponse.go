package helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/mahirjain_10/stock-alert-app/backend/internal/types"
)

func SendResponse(c *gin.Context, statusCode int, message string, data map[string]interface{}, errors map[string]string, success bool) {
	response := types.Response{
		Status:  statusCode,
		Message: message,
		Data:    data,
		Errors:  errors,
		Success: success,
	}

	c.JSON(statusCode, response)
}
