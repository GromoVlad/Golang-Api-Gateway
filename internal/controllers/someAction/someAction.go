package someAction

import (
	"gin_tonic/internal/support/localContext"
	"github.com/gin-gonic/gin"
)

// Endpoint - Какое-то действие
// SomeAction godoc
// @Summary      Какое-то действие
// @Tags         SomeAction
// @Security 	 BearerToken
// @Produce      json
// @Router       /some-action [post]
func Endpoint(ginContext *gin.Context) {
	context := localContext.LocalContext{Context: ginContext}
	context.StatusOK(gin.H{"data": "Удачное завершение", "success": true})
}
