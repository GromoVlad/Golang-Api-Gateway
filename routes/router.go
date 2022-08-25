package routes

import (
	"fmt"
	"gin_tonic/internal/controllers/example/login_json_controller"
	"gin_tonic/internal/controllers/example/post_form_data_controller"
	"gin_tonic/internal/middleware/logger_middleware"
	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.New()

	router.Use(logger_middleware.Logger())

	router.POST("/post", post_form_data_controller.PostFormData)
	router.POST("/loginJSON", login_json_controller.LoginJson)

	err := router.Run(":8080")
	if err != nil {
		fmt.Println("Произошла ошибка", err)
	}
}
