package router

import (
	"fuck-the-world/internal/controller"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.RouterGroup) {
	router.POST("/getUserInfo", controller.GetUserInfo)
}
