package utils

import (
	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, statusCode int, message string, err error) {
	if err != nil {
		c.JSON(statusCode, gin.H{
			"error":   message,
			"details": err.Error(),
		})
	} else {
		c.JSON(statusCode, gin.H{
			"error": message,
		})
	}
}
