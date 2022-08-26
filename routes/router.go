package routes

import (
	"fmt"
	"gin_tonic/internal/controllers/example/loginJson"
	"gin_tonic/internal/controllers/example/postFormData"
	"gin_tonic/internal/middleware/logger"
	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.New()

	router.Use(logger.Middleware())

	router.POST("/post", postFormData.Endpoint)
	router.POST("/loginJSON", loginJson.Endpoint)

	err := router.Run(":8080")
	if err != nil {
		fmt.Println("Произошла ошибка", err)
	}
}
