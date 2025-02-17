package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PingHandle(c *gin.Context) {
	result, err := c.Writer.WriteString("pong")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
	} else {
		c.JSON(200, gin.H{
			"message": result,
		})
	}
}
