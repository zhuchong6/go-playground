package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"minichat/docs"
	"minichat/service"
)

func Router() *gin.Engine {
	r := gin.Default()

	// swagger
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/index", service.GetIndex)

	userGroup := r.Group("/user")
	{
		userGroup.GET("/list", service.GetUserServiceList)
		userGroup.GET("/create", service.CreateUser)
		userGroup.GET("/delete", service.DeleteUser)
		userGroup.POST("/update", service.UpdateUser)
	}

	return r
}
