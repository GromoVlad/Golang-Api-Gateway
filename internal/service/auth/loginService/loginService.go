package loginService

import (
	"gin_tonic/internal/models/user"
	"gin_tonic/internal/repository/authRepository"
	"gin_tonic/internal/support/localContext"
	jwt4 "github.com/golang-jwt/jwt/v4"
	"os"
	"strconv"
	"time"
)

func Login(context localContext.LocalContext, user user.User) (string, string, int) {
	accessLifetimeM, _ := strconv.Atoi(os.Getenv("JWT_ACCESS_TOKEN_LIFETIME_M"))
	refreshLifetimeH, _ := strconv.Atoi(os.Getenv("JWT_REFRESH_TOKEN_LIFETIME_H"))

	accessToken := jwt4.NewWithClaims(jwt4.SigningMethodHS256, jwt4.MapClaims{
		"user_id": user.UserId,
		"name":    user.Name,
		"role_id": user.RoleId,
		"email":   user.Email,
		"exp":     jwt4.NewNumericDate(time.Now().Add(time.Minute * time.Duration(accessLifetimeM))),
	})
	accessTokenString, err := accessToken.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	context.InternalServerError(err)

	refreshToken := jwt4.NewWithClaims(jwt4.SigningMethodHS256, jwt4.MapClaims{
		"user_id": user.UserId,
		"name":    user.Name,
		"role_id": user.RoleId,
		"email":   user.Email,
		"exp":     jwt4.NewNumericDate(time.Now().Add(time.Hour * time.Duration(refreshLifetimeH))),
	})
	refreshTokenString, err := refreshToken.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	context.InternalServerError(err)

	authRepository.WriteRefreshJwtToken(context, refreshTokenString, user.UserId)

	return accessTokenString, refreshTokenString, accessLifetimeM
}
