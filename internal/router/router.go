package router

import (
	"fuck-the-world/config"
	"fuck-the-world/docs"
	"fuck-the-world/internal/controller"
	"fuck-the-world/internal/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title 接口文档
// @version 1.0
// @description 接口文档
// @termsOfService https://github.com/herzorf

// @contact.name herzorf
// @contact.url https://github.com/herzorf
// @contact.email 1446450047@qq.com

// @host 42.192.105.150:8888
// @BasePath "/api/v1"

func New() *gin.Engine {
	config.LoadConfigYaml()
	r := gin.New()
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // 允许前端的地址
		AllowMethods:     []string{"POST"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true, // 允许携带 Cookie
	}))
	r.POST("/api/v1/login", controller.Login)
	{
		//需要登陆的接口
		authGroup := r.Group("/api/v1")
		authGroup.Use(middleware.AuthMiddleware(), middleware.AuthAdminMiddleWare())
		authGroup.POST("/createOperator", controller.CreateOperator)
		authGroup.POST("/ping", controller.Ping)
		authGroup.POST("/sendEmail", controller.SendEmail)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

func RunServer() {
	r := New()
	err := r.Run("0.0.0.0:8888")
	if err != nil {
		panic(err)
	}
}
