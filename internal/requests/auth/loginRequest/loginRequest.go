package loginRequest

import (
	"gin_tonic/internal/support/localContext"
)

type Request struct {
	Email    string `form:"email"                json:"email"                binding:"required,email"`
	Password string `form:"password,required"    json:"password,required"    binding:"required,alphanum"`
}

func GetRequest(context localContext.LocalContext) Request {
	var request Request
	err := context.Context.ShouldBindJSON(&request)
	context.BadRequestError(err)
	return request
}
