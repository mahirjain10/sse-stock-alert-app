package helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/mahirjain_10/stock-alert-app/backend/internal/types"
	// "github.com/mahirjain_10/stock-alert-app/backend/internal/types"
)

// SendResponse sends a standardized JSON response with the given status code, message, data, and error details
func SendResponse(c *gin.Context, statusCode int, message string, data map[string]interface{}, errors map[string]string, success bool) {
	// Create response struct
	response := types.Response{
		Status:  statusCode,
		Message: message,
		Data:    data,
		Errors:  errors,
		Success: success,
	}

	// Send the response as JSON
	c.JSON(statusCode, response)
}
