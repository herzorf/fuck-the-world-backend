package router

import (
	"bookkeeping-server/internal/controller"
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", controller.PingHandle)
	return r
}
