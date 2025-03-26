package controller

import (
	"fuck-the-world/unit"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ping(c *gin.Context) {
	unit.RespondJSON(c, http.StatusOK, "pong", nil)
}
