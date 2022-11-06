package routes

import (
	"fmt"
	"gin_tonic/docs"
	"gin_tonic/internal/controllers/auth/login"
	"gin_tonic/internal/controllers/auth/refreshToken"
	"gin_tonic/internal/controllers/someAction"
	"gin_tonic/internal/controllers/user/deleteUser"
	"gin_tonic/internal/controllers/user/listUser"
	"gin_tonic/internal/controllers/user/registrationUser"
	"gin_tonic/internal/controllers/user/updateUser"
	"gin_tonic/internal/middleware/authSupportRole"
	"gin_tonic/internal/middleware/authWaiterRole"
	"gin_tonic/internal/middleware/globalLoggerMiddleware"
	"gin_tonic/internal/middleware/userGroupMiddleware"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/files"
	swaggerFiles "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag"
	"os"
)

func Run() {
	router := gin.New()

	/** Глобальные middleware */
	router.Use(globalLoggerMiddleware.Middleware())
	/** Восстановление после ошибки */
	router.Use(gin.Recovery())

	/** Роуты */
	api(router)

	/** Документация проекта */
	swaggerInfo(docs.SwaggerInfo)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := router.Run(":8081")
	if err != nil {
		fmt.Println("Произошла ошибка", err)
	}
}

func api(router *gin.Engine) {
	/** Роуты */
	router.GET("user/list", authSupportRole.Middleware(), listUser.Endpoint)

	userGroup := router.Group("/user")
	userGroup.Use(userGroupMiddleware.Middleware())
	{
		userGroup.POST("/registration", registrationUser.Endpoint)
		userGroup.PUT("/:userId", updateUser.Endpoint)
		userGroup.DELETE("/:userId", deleteUser.Endpoint)
	}

	authGroup := router.Group("/auth")
	{
		authGroup.POST("/login", login.Endpoint)
		authGroup.POST("/refresh-token", refreshToken.Endpoint)
	}

	router.POST("/some-action", authWaiterRole.Middleware(), someAction.Endpoint)
}

func swaggerInfo(swaggerInfo *swag.Spec) {
	swaggerInfo.Title = os.Getenv("PROJECT_TITLE")
	swaggerInfo.Description = os.Getenv("PROJECT_DESCRIPTION")
	swaggerInfo.Version = os.Getenv("PROJECT_VERSION")
	swaggerInfo.Host = os.Getenv("PROJECT_HOST")
	swaggerInfo.BasePath = os.Getenv("PROJECT_BASE_PATH")
	swaggerInfo.Schemes = []string{"http"}
}
