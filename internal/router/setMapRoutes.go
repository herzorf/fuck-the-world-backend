package router

import (
	"fuck-the-world/internal/controller"
	"github.com/gin-gonic/gin"
)

func setMapRoutes(router *gin.RouterGroup) {
	router.POST("createMap", controller.CreateMap)
}
