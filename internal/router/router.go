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
	docs.SwaggerInfo.BasePath = "/api"
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // 允许前端的地址
		AllowMethods:     []string{"POST"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true, // 允许携带 Cookie
	}))
	r.POST("/api/login", controller.Login)
	r.GET("/ping", controller.Ping)

	//需要登陆的接口
	authGroup := r.Group("/api")
	{
		authGroup.Use(middleware.AuthMiddleware())
		setMapRoutes(authGroup.Group("/map"))
		authGroup.POST("/sendEmail", controller.SendEmail)
		//用户相关接口
		SetupUserRoutes(authGroup.Group("/user"))
	}
	//需要登录和管理员权限的接口
	authAdminGroup := r.Group("/api")
	{
		authAdminGroup.Use(middleware.AuthMiddleware(), middleware.AuthAdminMiddleWare())
		//操作员相关接口
		SetupOperatorRoutes(authAdminGroup.Group("/operator"))
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
