package helpers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BindAndValidateJSON(c *gin.Context, input interface{}) bool {
	if err := c.ShouldBindJSON(input); err != nil {
		fmt.Printf("Error while decoding JSON: %v\n", err)
		SendResponse(c,http.StatusBadRequest,"Failed to decode JSON",nil,nil,false)
		return false
	}
	return true
}
