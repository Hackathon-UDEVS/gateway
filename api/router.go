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

	client := router.Group("/client")

	{
		client.POST("/create-tender", handler.CreateTender)
		client.PUT("/update-tender", handler.UpdateTender)
		client.DELETE("/delete-tender/:id", handler.DeleteTender)
		client.GET("/getAll-tenders", handler.GetTenders)
		client.GET("/tenders/sort", handler.SortTenders)
	}

	contractor := router.Group("/contractor")
	{
		contractor.POST("/submit-bid", handler.SubmitBid)
		contractor.GET("/bids", handler.GetListOfBids)
	}

	user := router.Group("/user")
	{
		user.POST("/login", handler.Login)
		user.POST("/register", handler.Register)
		user.POST("/verify-email", handler.VerifyEmail)
		user.GET("/get-user/:id", handler.GetUserByID)
		user.GET("getAll-users", handler.GetAllUsers)
		user.PUT("update-user", handler.UpdateUser)
	}

	return router
}
