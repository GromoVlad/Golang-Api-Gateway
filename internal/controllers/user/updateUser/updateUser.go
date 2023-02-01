package updateUser

import (
	"gin_tonic/internal/repository/userRepository"
	"gin_tonic/internal/requests/user/updateUserRequest"
	"gin_tonic/internal/response/baseResponse"
	"gin_tonic/internal/support/localContext"
	"github.com/gin-gonic/gin"
)

// Endpoint - Обновить данные пользователя
// UpdateUser godoc
// @Summary      Обновить данные пользователя
// @Tags         Users
// @Produce      json
// @Param        userId  path  int  true  "Идентификатор пользователя"
// @Param  		 RequestBody  body  updateUserRequest.Request  true	"Тело запроса"
/** @Success      200  {object}  baseResponse.BaseResponse{data=baseResponse.Response} "desc" */
// @Router       /user/{userId} [put]
func Endpoint(ginContext *gin.Context) {
	context := localContext.LocalContext{Context: ginContext}
	request := updateUserRequest.GetRequest(context)
	userRepository.UpdateUser(context, request)

	data := baseResponse.Response{Status: "Пользователь обновлен"}
	result := baseResponse.BaseResponse{Data: data, Success: true}
	context.StatusOK(gin.H{"data": result.Data, "success": result.Success})
}
