package routes

import (
	"fmt"
	"gin_tonic/docs"
	"gin_tonic/internal/controllers/apiGateway/book/createBook"
	"gin_tonic/internal/controllers/apiGateway/book/deleteBook"
	"gin_tonic/internal/controllers/apiGateway/book/findBook"
	"gin_tonic/internal/controllers/apiGateway/book/listBook"
	"gin_tonic/internal/controllers/apiGateway/book/updateBook"
	"gin_tonic/internal/controllers/auth/login"
	"gin_tonic/internal/controllers/auth/refreshToken"
	"gin_tonic/internal/controllers/user/deleteUser"
	"gin_tonic/internal/controllers/user/listUser"
	"gin_tonic/internal/controllers/user/registrationUser"
	"gin_tonic/internal/controllers/user/updateUser"
	"gin_tonic/internal/middleware/authSupportRole"
	"gin_tonic/internal/middleware/globalLoggerMiddleware"
	"gin_tonic/internal/middleware/userGroupMiddleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/swaggo/files"
	swaggerFiles "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag"
	"log"
	"os"
)

func Run() {
	router := gin.New()

	/** Глобальные middleware */
	router.Use(globalLoggerMiddleware.Middleware())
	/** Восстановление после ошибки */
	router.Use(gin.Recovery())
	/** Подгружаем данные из .env */
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Ошибка загрузки переменных из .env: %s", err.Error())
	}

	/** Роуты */
	api(router)

	/** Документация проекта */
	swaggerInfo(docs.SwaggerInfo)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := router.Run(":8082")
	if err != nil {
		fmt.Println("Произошла ошибка", err)
	}
}

func api(router *gin.Engine) {
	/** Роуты */
	router.GET("user/list", authSupportRole.Middleware(), listUser.Endpoint)
	//router.GET("user/list", listUser.Endpoint)

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

	// router.POST("/some-action", authWaiterRole.Middleware(), someAction.Endpoint)

	apiGatewayBook := router.Group("/api-gateway/book")
	{
		apiGatewayBook.GET("/:bookId", findBook.Endpoint)
		apiGatewayBook.GET("/list", listBook.Endpoint)
		apiGatewayBook.POST("/", createBook.Endpoint)
		apiGatewayBook.PUT("/:bookId", updateBook.Endpoint)
		apiGatewayBook.DELETE("/:bookId", deleteBook.Endpoint)
	}
}

func swaggerInfo(swaggerInfo *swag.Spec) {
	swaggerInfo.Title = os.Getenv("PROJECT_TITLE")
	swaggerInfo.Description = os.Getenv("PROJECT_DESCRIPTION")
	swaggerInfo.Version = os.Getenv("PROJECT_VERSION")
	swaggerInfo.Host = os.Getenv("PROJECT_HOST")
	swaggerInfo.BasePath = os.Getenv("PROJECT_BASE_PATH")
	swaggerInfo.Schemes = []string{"http"}
}
