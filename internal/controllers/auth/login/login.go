package login

import (
	"gin_tonic/internal/repository/userRepository"
	"gin_tonic/internal/requests/auth/loginRequest"
	"gin_tonic/internal/response/auth/loginResponse"
	"gin_tonic/internal/service/auth/loginService"
	"gin_tonic/internal/service/hash/passwordService"
	"gin_tonic/internal/support/localContext"
	"github.com/gin-gonic/gin"
	"time"
)

// Endpoint - Логин
// LoginUser godoc
// @Summary      Логин
// @Tags         Auth
// @Produce      json
// @Param  		 RequestBody  body  loginRequest.Request  true	"Тело запроса"
// @Success      200  {object}  loginResponse.BaseResponse{data=loginResponse.Response} "desc"
// @Router       /auth/login [post]
func Endpoint(ginContext *gin.Context) {
	context := localContext.LocalContext{Context: ginContext}
	request := loginRequest.GetRequest(context)

	user := userRepository.FindOrFailByEmail(context, request.Email)
	password := passwordService.AppendSaltToPassword(request.Password)
	passwordService.CheckPassword(context, password, user.Password)
	accessTokenString, refreshTokenString, accessLifetimeM := loginService.Login(context, user)

	data := loginResponse.Response{
		AccessToken:     accessTokenString,
		RefreshToken:    refreshTokenString,
		AccessExpiresAt: time.Now().Add(time.Minute * time.Duration(accessLifetimeM)),
	}
	result := loginResponse.BaseResponse{Data: data, Success: true}
	context.StatusOK(gin.H{"data": result.Data, "success": result.Success})
}
