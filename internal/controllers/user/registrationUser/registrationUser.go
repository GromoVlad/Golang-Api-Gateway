package registrationUser

import (
	"gin_tonic/internal/repository/userRepository"
	"gin_tonic/internal/requests/user/createUserRequest"
	"gin_tonic/internal/service/hash/passwordService"
	"gin_tonic/internal/support/localContext"
	"github.com/gin-gonic/gin"
)

// Endpoint - Создать пользователя
// RegistrationUser godoc
// @Summary      Регистрация пользователя
// @Tags         Users
// @Produce      json
// @Param  		 RequestBody  body  createUserRequest.DTO  true	"Тело запроса"
// @Success      201  {object}  BaseResponse
// @Router       /user/registration [post]
func Endpoint(ginContext *gin.Context) {
	context := localContext.LocalContext{Context: ginContext}
	dto := createUserRequest.GetRequest(context)
	dto.Password = passwordService.GetPasswordHash(context, dto.Password)

	userRepository.CreateUser(context, dto)

	result := BaseResponse{Data: Response{Status: "Пользователь создан"}, Success: true}
	context.StatusCreated(gin.H{"data": result.Data, "success": result.Success})
}

type BaseResponse struct {
	Data    Response `json:"data"`
	Success bool     `json:"success"`
}

type Response struct {
	Status string `json:"status"`
}
