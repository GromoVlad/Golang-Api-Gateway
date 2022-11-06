package jwtToken

import (
	"time"
)

type JwtToken struct {
	JwtTokenId   int       `db:"jwt_token_id" json:"jwt_token_id" example:"42" format:"int"`
	RefreshToken string    `db:"refresh_token" json:"refresh_token" format:"string"`
	UserId       int       `db:"user_id" json:"user_id" example:"42" format:"int"`
	CreatedAt    time.Time `db:"created_at" json:"created_at" example:"2022-01-01 00:00:00" format:"string"`
}
