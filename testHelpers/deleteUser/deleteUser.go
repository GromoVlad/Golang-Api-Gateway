package deleteUser

import (
	"gin_tonic/internal/repository/authRepository"
	"gin_tonic/internal/repository/userRepository"
	"gin_tonic/internal/support/localContext"
	"gin_tonic/testHelpers/findUser"
)

func DeleteUser(context localContext.LocalContext) {
	UserId := findUser.FindUser(context).UserId
	authRepository.DeleteTokens(context, UserId)
	userRepository.DeleteUser(context, UserId)
}
