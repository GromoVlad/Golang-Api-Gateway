package findUser

import (
	"gin_tonic/internal/models/user"
	"gin_tonic/internal/repository/userRepository"
	"gin_tonic/internal/support/localContext"
	"gin_tonic/testHelpers/createUser"
)

func FindUser(context localContext.LocalContext) user.User {
	return userRepository.FindOrFailByEmail(context, createUser.EMAIL)
}
