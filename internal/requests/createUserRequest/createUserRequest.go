package createUserRequest

import (
	"gin_tonic/internal/support/localContext"
)

type Request struct {
	Name     string `form:"name"                 json:"name"                 binding:"required"`
	Email    string `form:"email"                json:"email"                binding:"required,email"`
	RoleId   int    `form:"role_id"              json:"role_id"              binding:"required,number"`
	Phone    string `form:"phone,omitempty"      json:"phone,omitempty"      binding:"omitempty,numeric"`
	Password string `form:"password,omitempty"   json:"password,omitempty"   binding:"omitempty,alphanum"`
	HoReCaId int    `form:"horeca_id,omitempty"  json:"horeca_id,omitempty"  binding:"omitempty,number"`
}

func GetRequest(context localContext.LocalContext) Request {
	var request Request
	err := context.Context.ShouldBindJSON(&request)
	context.CheckBadRequestError(err)
	return request
}
