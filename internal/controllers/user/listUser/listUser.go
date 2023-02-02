package listUser

import (
	"gin_tonic/internal/models/user"
	"gin_tonic/internal/repository/userRepository"
	"gin_tonic/internal/requests/user/listUserRequest"
	"gin_tonic/internal/support/localContext"
	"github.com/gin-gonic/gin"
)

// Endpoint - Пагинированный список пользователей
// ListUser godoc
// @Summary      Пагинированный список пользователей
// @Tags         Users
// @Produce      json
// @Security 	 BearerToken
// @Param		 Authorization	header	string	true  "Добавить 'Bearer' перед токеном доступа"
// @Param  		 page   query	int	 	false	"Номер страницы"  minimum(1)
// @Param  		 limit  query	int	 	false	"Кол-во записей на странице" minimum(1)	maximum(20)
// @Param  		 search  query	string	false	"Поиск по имени"
// @Success      200  {object}  BaseResponse
// @Router       /user/list [get]
func Endpoint(ginContext *gin.Context) {
	context := localContext.LocalContext{Context: ginContext}
	request := listUserRequest.GetRequest(context)
	users, totalPage := userRepository.FindUsers(context, request)

	data := ListUserResponse{
		Users:       users,
		CurrentPage: request.Page,
		Limit:       request.Limit,
		TotalPage:   totalPage,
	}
	result := BaseResponse{Data: data, Success: true}
	context.StatusOK(gin.H{"data": result.Data, "success": result.Success})
}

type BaseResponse struct {
	Data    ListUserResponse `json:"data"`
	Success bool             `json:"success"`
}

type ListUserResponse struct {
	CurrentPage int         `json:"current_page"  example:"1"`
	Limit       int         `json:"limit" example:"10"`
	TotalPage   int         `json:"total_page" example:"100"`
	Users       []user.User `json:"users"`
}
