package router

import (
	_ "mini-admin/docs"
	"mini-admin/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()

	//add swagger
	// docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	indexGroup := r.Group("/index")
	indexGroup.GET("", service.IndexService)

	userGroup := r.Group("/users")
	userGroup.GET("/getUserByName", service.GetUserByName)

	return r
}
