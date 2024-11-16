package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"gateway/api/handler"
)

// @title Test API
// @version 1.0
// @description Test API with Swagger documentation
// @host localhost:8080
// @BasePath /api/v1

func InitRouter(handler *handler.Handler) *gin.Engine {
	router := gin.Default()

	router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Adjust for your specific origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// enforcer, err := casbin.NewEnforcer("./casbin/model.conf", "./casbin/policy.csv")
	// if err != nil {
	// 	panic(err)
	// }
	// router.Use(middlerware.NewAuth(enforcer))

	test := router.Group("test/")

	{
		test.GET("test1", handler.Test)
	}

	return router
}
