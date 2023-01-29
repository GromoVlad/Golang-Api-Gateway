package someAction

import (
	"gin_tonic/internal/database/DB"
	"gin_tonic/internal/support/localContext"
	"github.com/gin-gonic/gin"
)

// Endpoint - Какое-то действие
// SomeAction godoc
// @Summary      Какое-то действие
// @Tags         SomeAction
// @Produce      json
// @Router       /some-action [post]
func Endpoint(ginContext *gin.Context) {
	context := localContext.LocalContext{Context: ginContext}
	DB.Connect()
	context.StatusOK(gin.H{"success": true})
}
