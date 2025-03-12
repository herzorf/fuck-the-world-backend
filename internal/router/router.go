package router

import (
	"bookkeeping-server/docs"
	"bookkeeping-server/internal/controller"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/

//	@securityDefinitions.basic	BasicAuth

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/

func New() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", controller.PingHandle)
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
