package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	handler "gateway/api/handler"
)

// @title Test API
// @version 1.0
// @description Test API with Swagger documentation
// @host localhost:8080
// @BasePath /api/v1

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	h := handler.NewHandler()

	// Swagger docs
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api/v1")
	{
		api.GET("/test", h.Test)
	}

	return router
}
