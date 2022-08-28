package createUser

import (
	"gin_tonic/internal/repository/userRepository"
	"gin_tonic/internal/requests/createUserRequest"
	"github.com/gin-gonic/gin"
	"net/http"
)

var Endpoint = func(context *gin.Context) {
	request, err := createUserRequest.GetRequest(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = userRepository.CreateUser(request)
	if err != nil {
		context.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"status": "User created"})
}
