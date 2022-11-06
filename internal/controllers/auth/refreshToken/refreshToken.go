package refreshToken

import (
	"gin_tonic/internal/repository/authRepository"
	"gin_tonic/internal/repository/userRepository"
	"gin_tonic/internal/requests/auth/refreshTokenRequest"
	"gin_tonic/internal/response/auth/loginResponse"
	"gin_tonic/internal/service/auth/loginService"
	"gin_tonic/internal/support/localContext"
	"github.com/gin-gonic/gin"
	"time"
)

// Endpoint - Обновить access токен по refresh токену
// RefreshToken godoc
// @Summary      Обновить access токен по refresh токену
// @Tags         Auth
// @Produce      json
// @Param  		 RequestBody  body  refreshTokenRequest.Request  true	"Тело запроса"
// @Success      200  {object}  loginResponse.BaseResponse{data=loginResponse.Response} "desc"
// @Router       /auth/refresh-token [post]
func Endpoint(ginContext *gin.Context) {
	context := localContext.LocalContext{Context: ginContext}
	request := refreshTokenRequest.GetRequest(context)

	refreshToken := authRepository.FindOrFailToken(context, request.RefreshToken)
	user := userRepository.FindUser(context, refreshToken.UserId)
	accessTokenString, refreshTokenString, accessLifetimeM := loginService.Login(context, user)

	data := loginResponse.Response{
		AccessToken:     accessTokenString,
		RefreshToken:    refreshTokenString,
		AccessExpiresAt: time.Now().Add(time.Minute * time.Duration(accessLifetimeM)),
	}
	result := loginResponse.BaseResponse{Data: data, Success: true}
	context.StatusOK(gin.H{"data": result.Data, "success": result.Success})
}
