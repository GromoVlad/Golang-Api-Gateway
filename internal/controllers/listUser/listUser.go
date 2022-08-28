package listUser

import (
	"gin_tonic/internal/repository/userRepository"
	"github.com/gin-gonic/gin"
	"net/http"
)

var Endpoint = func(context *gin.Context) {
	users, err := userRepository.FindAllUser()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"users": users})
}
