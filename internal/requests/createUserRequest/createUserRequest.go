package createUserRequest

import "github.com/gin-gonic/gin"

type Request struct {
	Name     string `form:"name"                 json:"name"                 binding:"required"`
	Email    string `form:"email"                json:"email"                binding:"required,email"`
	RoleId   int    `form:"role_id"              json:"role_id"              binding:"required,number"`
	Phone    string `form:"phone,omitempty"      json:"phone,omitempty"      binding:"omitempty,numeric"`
	Password string `form:"password,omitempty"   json:"password,omitempty"   binding:"omitempty,alphanum"`
	HoReCaId int    `form:"horeca_id,omitempty"  json:"horeca_id,omitempty"  binding:"omitempty,number"`
}

func GetRequest(context *gin.Context) (Request, error) {
	var request Request
	if err := context.ShouldBindJSON(&request); err != nil {
		return request, err
	}
	return request, nil
}
