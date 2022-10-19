package routes

import (
	"fmt"
	"gin_tonic/docs"
	"gin_tonic/internal/controllers/createUser"
	"gin_tonic/internal/controllers/deleteUser"
	"gin_tonic/internal/controllers/listUser"
	"gin_tonic/internal/controllers/updateUser"
	"gin_tonic/internal/middleware/globalLoggerMiddleware"
	"gin_tonic/internal/middleware/routeMiddleware"
	"gin_tonic/internal/middleware/userGroupMiddleware"
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
	router.Use(globalLoggerMiddleware.Middleware())

	/** Роуты */
	router.GET("user/list", routeMiddleware.Middleware(), listUser.Endpoint)

	userGroup := router.Group("/user")
	userGroup.Use(userGroupMiddleware.Middleware())
	{
		userGroup.POST("/", createUser.Endpoint)
		userGroup.PUT("/:userId", updateUser.Endpoint)
		userGroup.DELETE("/:userId", deleteUser.Endpoint)
	}

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
