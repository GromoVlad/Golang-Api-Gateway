package createUser

import (
	"gin_tonic/internal/repository/userRepository"
	"gin_tonic/internal/requests/createUserRequest"
	"gin_tonic/internal/response/baseResponse"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Endpoint - Создать пользователя
// CreateUser godoc
// @Summary      Создать пользователя
// @Tags         Users
// @Produce      json
// @Param  		 RequestBody  body  createUserRequest.Request  true	"Тело запроса"
// @Success      201  {object}  baseResponse.BaseResponse{data=Response} "desc"
// @Router       /create-user [post]
func Endpoint(context *gin.Context) {
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
	data := Response{Status: "Пользователь создан"}
	result := baseResponse.BaseResponse{Data: data, Success: true}
	context.JSON(http.StatusCreated, gin.H{"data": result.Data, "success": result.Success})
}

type Response struct {
	Status string `json:"status"`
}
