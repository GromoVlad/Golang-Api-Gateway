package updateUser

import (
	"gin_tonic/internal/repository/userRepository"
	"gin_tonic/internal/requests/updateUserRequest"
	"gin_tonic/internal/response/baseResponse"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Endpoint - Обновить данные пользователя
// UpdateUser godoc
// @Summary      Обновить данные пользователя
// @Tags         Users
// @Produce      json
// @Param        userId  path  int  true  "Идентификатор пользователя"
// @Param  		 RequestBody  body  updateUserRequest.Request  true	"Тело запроса"
// @Success      200  {object}  baseResponse.BaseResponse{data=Response} "desc"
// @Router       /update-user/{userId} [put]
func Endpoint(context *gin.Context) {
	request, err := updateUserRequest.GetRequest(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = userRepository.UpdateUser(request)
	if err != nil {
		context.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	data := Response{Status: "Пользователь обновлен"}
	result := baseResponse.BaseResponse{Data: data, Success: true}
	context.JSON(http.StatusOK, gin.H{"data": result.Data, "success": result.Success})
}

type Response struct {
	Status string `json:"status"`
}
