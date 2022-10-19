package userGroupMiddleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Middleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		log.Print("Груповой middleware на роуты user (create/update/delete)")
	}
}
