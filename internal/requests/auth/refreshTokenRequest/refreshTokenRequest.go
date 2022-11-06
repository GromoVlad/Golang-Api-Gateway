package refreshTokenRequest

import (
	"fmt"
	"gin_tonic/internal/support/localContext"
	jwt4 "github.com/golang-jwt/jwt/v4"
	"os"
)

type Request struct {
	RefreshToken string `form:"refresh_token"   json:"refresh_token"   binding:"required"`
}

func GetRequest(context localContext.LocalContext) Request {
	var request Request
	err := context.Context.ShouldBindJSON(&request)
	context.BadRequestError(err)

	_, err = jwt4.Parse(request.RefreshToken, func(token *jwt4.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt4.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("не определен алгоритм расшифровки токена: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	context.TokenError(err)

	return request
}
