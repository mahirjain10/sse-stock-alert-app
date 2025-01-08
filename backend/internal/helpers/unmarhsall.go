package helpers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// BindAndValidateJSON binds the incoming JSON to the provided input object and validates it
// Returns true if binding is successful, otherwise sends a bad request response and returns false.
func BindAndValidateJSON(c *gin.Context, input interface{}) bool {
	// Attempt to bind the JSON payload to the input object
	if err := c.ShouldBindJSON(input); err != nil {
		// Log the error for debugging purposes
		fmt.Printf("Error while decoding JSON: %v\n", err)
		
		// Send a bad request response with an error message
		SendResponse(c, http.StatusBadRequest, "Failed to decode JSON", nil, nil, false)
		
		// Return false as the binding failed
		return false
	}
	
	// Return true if binding was successful
	return true
}
