package deleteUser

import (
	"gin_tonic/internal/repository/userRepository"
	"gin_tonic/internal/support/localContext"
	"github.com/gin-gonic/gin"
	"strconv"
)

// Endpoint - Обновить данные пользователя
// DeleteUser godoc
// @Summary      Удалить пользователя
// @Tags         Users
// @Produce      json
// @Param        userId  path  int  true  "Идентификатор пользователя"
// @Success      200  {object}  BaseResponse
// @Router       /user/{userId} [delete]
func Endpoint(ginContext *gin.Context) {
	context := localContext.LocalContext{Context: ginContext}

	userId, err := strconv.Atoi(ginContext.Param("userId"))
	context.BadRequestError(err)

	userRepository.DeleteUser(context, userId)

	result := BaseResponse{Data: Response{Status: "Пользователь удален"}, Success: true}
	context.StatusOK(gin.H{"data": result.Data, "success": result.Success})
}

type BaseResponse struct {
	Data    Response `json:"data"`
	Success bool     `json:"success"`
}

type Response struct {
	Status string `json:"status"`
}
