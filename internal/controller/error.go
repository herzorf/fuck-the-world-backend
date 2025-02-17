package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func handleError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": err.Error(),
	})
}
