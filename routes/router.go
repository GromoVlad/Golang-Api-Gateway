package routes

import (
	"fmt"
	"gin_tonic/internal/controllers/createUser"
	"gin_tonic/internal/controllers/listUser"
	"gin_tonic/internal/controllers/updateUser"
	"gin_tonic/internal/middleware/logger"
	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.New()

	router.Use(logger.Middleware())

	router.POST("/list-user", listUser.Endpoint)
	router.POST("/create-user", createUser.Endpoint)
	router.PUT("/update-user/:userId", updateUser.Endpoint)

	err := router.Run(":8080")
	if err != nil {
		fmt.Println("Произошла ошибка", err)
	}
}
