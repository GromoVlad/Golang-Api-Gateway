package createUser

import (
	"gin_tonic/internal/repository/userRepository"
	"gin_tonic/internal/requests/createUserRequest"
	"gin_tonic/internal/response/baseResponse"
	"gin_tonic/internal/support/localContext"
	"github.com/gin-gonic/gin"
)

// Endpoint - Создать пользователя
// CreateUser godoc
// @Summary      Создать пользователя
// @Tags         Users
// @Produce      json
// @Param  		 RequestBody  body  createUserRequest.Request  true	"Тело запроса"
// @Success      201  {object}  baseResponse.BaseResponse{data=baseResponse.Response} "desc"
// @Router       /user [post]
func Endpoint(ginContext *gin.Context) {
	context := localContext.LocalContext{Context: ginContext}

	request := createUserRequest.GetRequest(context)

	userRepository.CreateUser(context, request)

	data := baseResponse.Response{Status: "Пользователь создан"}
	result := baseResponse.BaseResponse{Data: data, Success: true}
	context.SuccessStatusCreated(gin.H{"data": result.Data, "success": result.Success})
}
