package loginResponse

import (
	"time"
)

type BaseResponse struct {
	Data    any  `json:"data"`
	Success bool `json:"success"`
}

type Response struct {
	AccessToken     string    `json:"access_token"`
	RefreshToken    string    `json:"refresh_token"`
	AccessExpiresAt time.Time `json:"expires_at"`
}
