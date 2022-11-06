package authWaiterRole

import (
	"errors"
	"gin_tonic/internal/enums/role"
	"gin_tonic/internal/service/auth/bearerTokenService"
	"gin_tonic/internal/support/localContext"
	"github.com/gin-gonic/gin"
)

func Middleware() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		context := localContext.LocalContext{Context: ginContext}
		roleId := bearerTokenService.ParseRoleId(context)
		if roleId != role.WAITER {
			context.UnauthorizedError(errors.New("пользователь не является официантом, недостаточно прав"))
		}
	}
}
