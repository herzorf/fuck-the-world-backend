package controller

import (
	"fuck-the-world/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ping(c *gin.Context) {
	utils.RespondJSON(c, http.StatusOK, "pong", nil)
}
