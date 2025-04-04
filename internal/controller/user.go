package controller

import (
	"fuck-the-world/unit"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserInfo(c *gin.Context) {
	var role = c.MustGet("role")
	var username = c.MustGet("username")

	unit.RespondJSON(c, http.StatusOK, "", gin.H{
		"role":     role,
		"username": username,
	})
}
