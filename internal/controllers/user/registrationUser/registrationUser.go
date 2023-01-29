package registrationUser

import (
	"gin_tonic/internal/repository/userRepository"
	"gin_tonic/internal/requests/user/createUserRequest"
	"gin_tonic/internal/response/baseResponse"
	"gin_tonic/internal/service/hash/passwordService"
	"gin_tonic/internal/support/localContext"
	"github.com/gin-gonic/gin"
)

// Endpoint - Создать пользователя
// RegistrationUser godoc
// @Summary      Регистрация пользователя
// @Tags         Users
// @Produce      json
// @Param  		 RequestBody  body  createUserRequest.Request  true	"Тело запроса"
// @Success      201  {object}  baseResponse.BaseResponse{data=baseResponse.Response} "desc"
// @Router       /user/registration [post]
func Endpoint(ginContext *gin.Context) {
	context := localContext.LocalContext{Context: ginContext}
	dto := createUserRequest.GetRequest(context)
	dto.Password = passwordService.GetPasswordHash(context, dto.Password)

	userRepository.CreateUser(context, dto)

	data := baseResponse.Response{Status: "Пользователь создан"}
	result := baseResponse.BaseResponse{Data: data, Success: true}
	context.StatusCreated(gin.H{"data": result.Data, "success": result.Success})
}
