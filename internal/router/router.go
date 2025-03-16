package router

import (
	"bookkeeping-server/config"
	"bookkeeping-server/docs"
	"bookkeeping-server/internal/controller"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			项目文档
//	@version		1.0
//	@description	这是一个简单的记账项目

//	@contact.name	herzorf
//	@contact.url	https://github.com/herzorf
//	@contact.email	herzorf@icloud.com

//	@host		localhost:8080
//	@BasePath	/

func New() *gin.Engine {
	config.LoadConfigYaml()
	r := gin.Default()
	r.GET("/ping", controller.PingHandle)
	r.POST("/sendEmail", controller.SendEmail)

	docs.SwaggerInfo.Version = "1.0"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

func RunServer() {
	r := New()
	err := r.Run() // listen and serve on 0.0.0.0:8080
	if err != nil {
		panic(err)
	}
}
