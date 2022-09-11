package listUser

import (
	"gin_tonic/internal/repository/userRepository"
	"gin_tonic/internal/requests/listRepositoryRequest"
	"gin_tonic/internal/response/baseResponse"
	"gin_tonic/internal/response/listUserResponse"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Endpoint - Пагинированный список пользователей
// ListUser godoc
// @Summary      Пагинированный список пользователей
// @Tags         Users
// @Produce      json
// @Param  		 page   query	int	 	false	"Номер страницы"  minimum(1)
// @Param  		 limit  query	int	 	false	"Кол-во записей на странице" minimum(1)	maximum(20)
// @Param  		 search  query	string	false	"Поиск по имени"
// @Success      200  {object}  baseResponse.BaseResponse{data=listUserResponse.ListUserResponse} "desc"
// @Router       /list-user [get]
func Endpoint(context *gin.Context) {
	request, err := listRepositoryRequest.GetRequest(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	users, totalPage, err := userRepository.FindUsers(request)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	data := listUserResponse.ListUserResponse{
		Users:       users,
		CurrentPage: request.Page,
		Limit:       request.Limit,
		TotalPage:   totalPage,
	}
	result := baseResponse.BaseResponse{Data: data, Success: true}
	context.JSON(http.StatusOK, gin.H{"data": result.Data, "success": result.Success})
}
