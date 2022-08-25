package login_json_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type MyName struct {
	Name string `form:"name" json:"name" xml:"name"  binding:"required"`
}

var LoginJson = func(context *gin.Context) {
	var bodyStruct MyName
	if err := context.ShouldBindJSON(&bodyStruct); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if bodyStruct.Name != "Vlad" {
		context.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"status": "Hello, " + bodyStruct.Name + ". You are logged in"})
}
