package userRepository

import (
	"gin_tonic/internal/database/DB"
	"gin_tonic/internal/models/user"
	"log"
)

func FindAllUser() []user.User {
	var users []user.User
	err := DB.Connect().Select(&users, "SELECT * FROM users.users")
	if err != nil {
		log.Fatalln(err)
	}
	return users
}
