package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// PingHandle godoc
// @Summary ping
// @Description ping
// @Tags ping
// @Accept json
// @Produce json
// @Success 200
// @error 500
// @Router /ping [get]
func PingHandle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
