package router

import (
	"fuck-the-world/internal/controller"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.RouterGroup) {
	router.POST("/createOperator", controller.CreateOperator)
	router.POST("/deleteOperator", controller.DeleteOperator)
	router.POST("/queryOperatorList", controller.QueryOperatorList)
	router.POST("/updateOperator", controller.UpdateOperator)
}
