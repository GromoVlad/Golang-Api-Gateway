package createUserRequest

import (
	"gin_tonic/internal/support/localContext"
)

type Request struct {
	Name     string `form:"name"                 json:"name"                 binding:"required"`
	Email    string `form:"email"                json:"email"                binding:"required,email"`
	RoleId   int    `form:"role_id"              json:"role_id"              binding:"required,number"`
	Password string `form:"password"             json:"password"             binding:"required,alphanum"`
	Phone    string `form:"phone,omitempty"      json:"phone,omitempty"      binding:"omitempty,numeric"`
	VenueId  int    `form:"venue_id,omitempty"   json:"venue_id,omitempty"   binding:"omitempty,number"`
}

func GetRequest(context localContext.LocalContext) Request {
	var request Request
	err := context.Context.ShouldBindJSON(&request)
	context.BadRequestError(err)
	return request
}
