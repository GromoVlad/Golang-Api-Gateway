package routes

import (
	"fmt"
	"gin_tonic/docs"
	"gin_tonic/internal/controllers/createUser"
	"gin_tonic/internal/controllers/listUser"
	"gin_tonic/internal/controllers/updateUser"
	"gin_tonic/internal/middleware/logger"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/files"
	swaggerFiles "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag"
)

func Run() {
	router := gin.New()

	/** Глобальные middleware */
	router.Use(logger.Middleware())

	/** Роуты */
	router.GET("/list-user", listUser.Endpoint)
	router.POST("/create-user", createUser.Endpoint)
	router.PUT("/update-user/:userId", updateUser.Endpoint)

	/** Документация проекта */
	swaggerInfo(docs.SwaggerInfo)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := router.Run(":8081")
	if err != nil {
		fmt.Println("Произошла ошибка", err)
	}
}

func swaggerInfo(swaggerInfo *swag.Spec) {
	swaggerInfo.Title = "Gin-Tonic"
	swaggerInfo.Description = "Минифреймворк (референс - Laravel) из модулей Gin,Sqlx, Goose, Swaggo"
	swaggerInfo.Version = "1.0"
	swaggerInfo.Host = "localhost:8081"
	swaggerInfo.BasePath = "/"
	swaggerInfo.Schemes = []string{"http"}
}
