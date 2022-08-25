package post_form_data_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var PostFormData = func(context *gin.Context) {
	name := context.PostForm("name")
	context.JSON(http.StatusOK, gin.H{"status": "Hello, " + name + ". Request post_form_data_controller success"})
}
