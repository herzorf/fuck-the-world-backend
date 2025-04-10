package controller

import (
	"fuck-the-world/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserInfo(c *gin.Context) {
	var role = c.MustGet("role")
	var username = c.MustGet("username")

	utils.RespondJSON(c, http.StatusOK, "", gin.H{
		"role":     role,
		"username": username,
	})
}
