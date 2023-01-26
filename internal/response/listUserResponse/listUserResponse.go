package listUserResponse

import "gin_tonic/internal/models/user"

type ListUserResponse struct {
	CurrentPage int         `json:"current_page"  example:"1"`
	Limit       int         `json:"limit" example:"10"`
	TotalPage   int         `json:"total_page" example:"100"`
	Users       []user.User `json:"users"`
}

type BaseResponse struct {
	Data    ListUserResponse `json:"data"`
	Success bool             `json:"success"`
}
