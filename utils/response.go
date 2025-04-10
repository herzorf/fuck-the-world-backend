package utils

import "github.com/gin-gonic/gin"

func RespondJSON(c *gin.Context, code int, message string, result interface{}) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
		"result":  result,
	})
}
