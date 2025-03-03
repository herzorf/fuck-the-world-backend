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

func RunServer() {
	r := New()
	err := r.Run() // listen and serve on 0.0.0.0:8080
	if err != nil {
		panic(err)
	}
}
