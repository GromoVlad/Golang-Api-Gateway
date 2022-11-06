package passwordService

import (
	"fmt"
	"gin_tonic/internal/support/localContext"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

func GetPasswordHash(context localContext.LocalContext, password string) string {
	t := time.Now()
	password = AppendSaltToPassword(password)
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	context.BadRequestError(err)
	fmt.Println("\nGetPasswordHashEnd = ", time.Since(t))
	return string(bytes)
}

func CheckPassword(context localContext.LocalContext, password string, userPassword string) {
	t := time.Now()
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(password))
	context.PasswordError(err)
	fmt.Println("\nCheckPasswordEnd = ", time.Since(t))
}

func AppendSaltToPassword(password string) string {
	return os.Getenv("PASSWORD_PREFIX_SALT") + password + os.Getenv("PASSWORD_POSTFIX_SALT")
}
