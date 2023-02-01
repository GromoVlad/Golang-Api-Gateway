package createUser

import (
	"gin_tonic/internal/enums/role"
	"gin_tonic/internal/repository/userRepository"
	"gin_tonic/internal/requests/user/createUserRequest"
	"gin_tonic/internal/service/hash/passwordService"
	"gin_tonic/internal/support/localContext"
)

const USERNAME = "TestUser"
const PASSWORD = "12345678"
const EMAIL = "TestUser@yandex.ru"

func CreateUser(context localContext.LocalContext) {
	dto := createUserRequest.DTO{
		Name:     USERNAME,
		Email:    EMAIL,
		RoleId:   role.SUPPORT,
		Password: PASSWORD,
		Phone:    "89998887766",
		VenueId:  100,
	}

	password := passwordService.GetPasswordHash(context, dto.Password)
	dto.Password = password
	userRepository.CreateUser(context, dto)
}
