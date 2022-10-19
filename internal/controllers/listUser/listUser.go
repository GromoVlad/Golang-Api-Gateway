package listUser

import (
	"fmt"
	"gin_tonic/internal/repository/userRepository"
	"gin_tonic/internal/requests/listRepositoryRequest"
	"gin_tonic/internal/response/baseResponse"
	"gin_tonic/internal/response/listUserResponse"
	"gin_tonic/internal/support/context"
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
	response := context.Response{Context: ginContext}

	request, err := listRepositoryRequest.GetRequest(ginContext)
	response.CheckBadRequestError(err)

	users, totalPage, err := userRepository.FindUsers(request)
	response.CheckInternalServerError(err)

	err = logger.InfoLog("Список пользователей", fmt.Sprintf("%v", users))
	response.CheckInternalServerError(err)
	if response.Context.IsAborted() {
		return
	}

	data := listUserResponse.ListUserResponse{
		Users:       users,
		CurrentPage: request.Page,
		Limit:       request.Limit,
		TotalPage:   totalPage,
	}
	result := baseResponse.BaseResponse{Data: data, Success: true}
	response.SuccessStatusOK(gin.H{"data": result.Data, "success": result.Success})
}
