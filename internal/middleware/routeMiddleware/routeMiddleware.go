package routeMiddleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Middleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		log.Print("Отдельный middleware на роут list-user")
	}
}
