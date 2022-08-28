package updateUserRequest

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type Request struct {
	UserId   int
	Name     string `form:"name,omitempty"                  json:"name,omitempty"                  binding:"omitempty"`
	Email    string `form:"email,omitempty"                 json:"email,omitempty"                 binding:"omitempty,email"`
	RoleId   int    `form:"role_id,omitempty"               json:"role_id,omitempty"               binding:"omitempty,number"`
	Phone    string `form:"phone,omitempty"                 json:"phone,omitempty"                 binding:"omitempty,numeric"`
	Password string `form:"password,omitempty"              json:"password,omitempty"              binding:"omitempty,alphanum"`
	HoReCaId int    `form:"horeca_id,omitempty"             json:"horeca_id,omitempty"             binding:"omitempty,number"`
	Url      string `form:"password_recovery_url,omitempty" json:"password_recovery_url,omitempty" binding:"omitempty,uri"`
}

func GetRequest(context *gin.Context) (Request, error) {
	var request Request
	if err := context.ShouldBindJSON(&request); err != nil {
		return request, err
	}
	userId := context.Param("userId")
	request.UserId, _ = strconv.Atoi(userId)
	return request, nil
}
