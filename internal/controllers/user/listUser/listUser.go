package listUser

import (
	"fmt"
	"gin_tonic/internal/repository/userRepository"
	"gin_tonic/internal/requests/user/listRepositoryRequest"
	"gin_tonic/internal/response/baseResponse"
	"gin_tonic/internal/response/listUserResponse"
	"gin_tonic/internal/support/localContext"
	"gin_tonic/internal/support/logger"
	"github.com/gin-gonic/gin"
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
func Endpoint(ginContext *gin.Context) {
	context := localContext.LocalContext{Context: ginContext}
	request := listRepositoryRequest.GetRequest(context)
	users, totalPage := userRepository.FindUsers(context, request)
	logger.InfoLog(context, "Список пользователей", fmt.Sprintf("%v", users))

	data := listUserResponse.ListUserResponse{
		Users:       users,
		CurrentPage: request.Page,
		Limit:       request.Limit,
		TotalPage:   totalPage,
	}
	result := baseResponse.BaseResponse{Data: data, Success: true}
	context.StatusOK(gin.H{"data": result.Data, "success": result.Success})
}
