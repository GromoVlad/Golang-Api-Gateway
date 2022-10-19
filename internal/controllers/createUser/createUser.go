package createUser

import (
	"gin_tonic/internal/repository/userRepository"
	"gin_tonic/internal/requests/createUserRequest"
	"gin_tonic/internal/response/baseResponse"
	"gin_tonic/internal/support/context"
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
	response := context.Response{Context: ginContext}

	request, err := createUserRequest.GetRequest(ginContext)
	response.CheckBadRequestError(err)

	err = userRepository.CreateUser(request)
	response.CheckStatusConflictError(err)
	if response.Context.IsAborted() {
		return
	}

	data := baseResponse.Response{Status: "Пользователь создан"}
	result := baseResponse.BaseResponse{Data: data, Success: true}
	response.SuccessStatusCreated(gin.H{"data": result.Data, "success": result.Success})
}
