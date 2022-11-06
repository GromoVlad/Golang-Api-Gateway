package updateUserRequest

import (
	"gin_tonic/internal/support/localContext"
	"strconv"
)

type Request struct {
	UserId   int
	Name     string `form:"name,omitempty"                  json:"name,omitempty"                  binding:"omitempty"`
	Email    string `form:"email,omitempty"                 json:"email,omitempty"                 binding:"omitempty,email"`
	RoleId   int    `form:"role_id,omitempty"               json:"role_id,omitempty"               binding:"omitempty,number"`
	Phone    string `form:"phone,omitempty"                 json:"phone,omitempty"                 binding:"omitempty,numeric"`
	Password string `form:"password,omitempty"              json:"password,omitempty"              binding:"omitempty,alphanum"`
	VenueId  int    `form:"venue_id,omitempty"             json:"venue_id,omitempty"             binding:"omitempty,number"`
	Url      string `form:"password_recovery_url,omitempty" json:"password_recovery_url,omitempty" binding:"omitempty,uri"`
}

func GetRequest(context localContext.LocalContext) Request {
	var request Request
	err := context.Context.ShouldBindJSON(&request)
	context.BadRequestError(err)

	userId := context.Context.Param("userId")
	request.UserId, err = strconv.Atoi(userId)
	context.BadRequestError(err)

	return request
}
