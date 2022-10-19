package deleteUserRequest

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type Request struct {
	UserId int
}

func GetRequest(context *gin.Context) (Request, error) {
	var request Request
	userId := context.Param("userId")
	request.UserId, _ = strconv.Atoi(userId)
	return request, nil
}
