package router

import (
	"fuck-the-world/internal/controller"
	"github.com/gin-gonic/gin"
)

func setMapRoutes(router *gin.RouterGroup) {
	router.POST("createMap", controller.CreateMap)
	router.POST("queryMapList", controller.QueryMapList)
	router.POST("deleteMap", controller.DeleteMap)
	router.POST("updateMap", controller.UpdateMap)
}
