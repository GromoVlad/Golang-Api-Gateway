package authRepository

import (
	"errors"
	"fmt"
	"gin_tonic/internal/database/DB"
	"gin_tonic/internal/models/jwtToken"
	"gin_tonic/internal/support/localContext"
	"time"
)

func FindOrFailToken(context localContext.LocalContext, token string) jwtToken.JwtToken {
	var refreshToken jwtToken.JwtToken
	_ = DB.Connect().Get(&refreshToken, "SELECT * FROM auth.jwt_tokens WHERE refresh_token = $1", token)
	if refreshToken.JwtTokenId == 0 {
		context.NotFoundError(errors.New(fmt.Sprintf("Переданный токен не найден")))
	}
	return refreshToken
}

func WriteRefreshJwtToken(context localContext.LocalContext, token string, userId int) {
	transaction := DB.Connect().MustBegin()

	_, err := transaction.NamedExec(
		"DELETE FROM auth.jwt_tokens WHERE user_id = :user_id",
		&jwtToken.JwtToken{UserId: userId},
	)
	context.StatusConflictError(err)

	_, err = transaction.NamedExec(
		"INSERT INTO auth.jwt_tokens (refresh_token, user_id, created_at) VALUES (:refresh_token, :user_id, :created_at)",
		&jwtToken.JwtToken{
			RefreshToken: token,
			UserId:       userId,
			CreatedAt:    time.Now(),
		},
	)
	context.StatusConflictError(err)

	err = transaction.Commit()
	context.InternalServerError(err)
}
