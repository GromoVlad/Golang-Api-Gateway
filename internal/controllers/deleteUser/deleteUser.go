package deleteUser

import (
	"gin_tonic/internal/repository/userRepository"
	"gin_tonic/internal/response/baseResponse"
	"gin_tonic/internal/support/context"
	"github.com/gin-gonic/gin"
	"strconv"
)

// Endpoint - Обновить данные пользователя
// DeleteUser godoc
// @Summary      Удалить пользователя
// @Tags         Users
// @Produce      json
// @Param        userId  path  int  true  "Идентификатор пользователя"
// @Success      200  {object}  baseResponse.BaseResponse{data=baseResponse.Response} "desc"
// @Router       /user/{userId} [delete]
func Endpoint(ginContext *gin.Context) {
	response := context.Response{Context: ginContext}

	userId, err := strconv.Atoi(ginContext.Param("userId"))
	response.CheckBadRequestError(err)

	err = userRepository.DeleteUser(userId)
	response.CheckStatusConflictError(err)
	if response.Context.IsAborted() {
		return
	}

	data := baseResponse.Response{Status: "Пользователь удален"}
	result := baseResponse.BaseResponse{Data: data, Success: true}
	response.SuccessStatusOK(gin.H{"data": result.Data, "success": result.Success})
}
