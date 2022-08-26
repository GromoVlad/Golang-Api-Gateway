package logger

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Middleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		t := time.Now()
		// Set example variable
		context.Set("example", "12345")
		// перед запросом
		context.Next()
		// после запроса
		latency := time.Since(t)
		log.Print("\nlatency = ", latency)
		// доступ к статусу, который отправляем
		status := context.Writer.Status()
		log.Print("status = ", status)
	}
}
