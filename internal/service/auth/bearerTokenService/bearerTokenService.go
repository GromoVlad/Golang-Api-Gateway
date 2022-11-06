package bearerTokenService

import (
	"errors"
	"fmt"
	"gin_tonic/internal/support/headers"
	"gin_tonic/internal/support/localContext"
	jwt4 "github.com/golang-jwt/jwt/v4"
	"os"
	"strings"
)

func ParseRoleId(context localContext.LocalContext) int {
	header := headers.Headers{}
	if err := context.Context.ShouldBindHeader(&header); err != nil {
		context.BadRequestError(err)
	}

	tokenString := header.BearerToken
	contain := strings.Contains(tokenString, "Bearer ")
	if !contain {
		context.BadRequestError(errors.New("некорректный токен авторизации, ожидается BearerToken"))
	}

	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	token, err := jwt4.Parse(tokenString, func(token *jwt4.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt4.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("не определен алгоритм расшифровки токена: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	context.TokenError(err)

	claims, ok := token.Claims.(jwt4.MapClaims)
	if !(ok && token.Valid) {
		context.BadRequestError(errors.New("передан невалидный токен авторизации"))
	}

	floatRoleId, ok := claims["role_id"].(float64)
	if !ok {
		context.BadRequestError(errors.New("передан невалидный токен авторизации"))
	}

	return int(floatRoleId)
}
