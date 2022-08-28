package userRepository

import (
	"database/sql"
	"errors"
	"fmt"
	"gin_tonic/internal/database/DB"
	"gin_tonic/internal/enums"
	"gin_tonic/internal/models/user"
	"gin_tonic/internal/requests/createUserRequest"
	"gin_tonic/internal/requests/updateUserRequest"
	"time"
)

func FindUser(userId int) (user.User, error) {
	var findUser user.User
	err := DB.Connect().Get(&findUser, "SELECT * FROM users.users WHERE user_id = $1", userId)
	if err != nil {
		return findUser, err
	}
	if findUser.UserId == 0 {
		err = errors.New(fmt.Sprintf(
			"Пользователь с идентификатором %d не зарегистрирован в системе",
			userId,
		))
		return findUser, err
	}
	return findUser, nil
}

func FindAllUser() ([]user.User, error) {
	var users []user.User
	err := DB.Connect().Select(&users, "SELECT * FROM users.users")
	if err != nil {
		return users, err
	}
	return users, nil
}

func CreateUser(request createUserRequest.Request) error {
	var findUser user.User

	if request.Email != "" {
		err := DB.Connect().Get(&findUser, "SELECT user_id FROM users.users WHERE email = $1", request.Email)
		if err != nil {
			return err
		}
		if findUser.UserId != 0 {
			err = errors.New("Пользователь с email " + request.Email + " уже зарегистрирован в системе")
			return err
		}
	}

	transaction := DB.Connect().MustBegin()
	_, err := transaction.NamedExec(
		"INSERT INTO users.users (name, role_id, phone, password, email, horeca_id, password_recovery_url, messenger, created_at) "+
			"VALUES (:name, :role_id, :phone, :password, :email, :horeca_id, :password_recovery_url, :messenger, :created_at)",
		&user.User{
			Name:                request.Name,
			RoleId:              request.RoleId,
			Phone:               sql.NullString{String: request.Phone, Valid: request.Phone != ""},
			Password:            sql.NullString{String: request.Password, Valid: request.Password != ""},
			Email:               sql.NullString{String: request.Email, Valid: request.Email != ""},
			HoReCaId:            sql.NullInt16{Int16: int16(request.HoReCaId), Valid: request.HoReCaId != 0},
			PasswordRecoveryUrl: sql.NullString{},
			Messenger:           sql.NullString{String: enums.TELEGRAM, Valid: true},
			CreatedAt:           sql.NullTime{Time: time.Now(), Valid: true},
			UpdatedAt:           sql.NullTime{},
		},
	)

	if err != nil {
		return err
	}
	err = transaction.Commit()
	if err != nil {
		return err
	}
	return nil
}

func UpdateUser(request updateUserRequest.Request) error {

	findUser, err := FindUser(request.UserId)
	if err != nil {
		return err
	}
	mappingUser(&findUser, request)
	transaction := DB.Connect().MustBegin()
	_, err = transaction.NamedExec(
		"UPDATE users.users SET updated_at = :updated_at, name = :name, role_id = :role_id, "+
			"phone = :phone, password = :password, email = :email, horeca_id = :horeca_id, "+
			"password_recovery_url = :password_recovery_url WHERE user_id = :user_id",
		&findUser,
	)
	if err != nil {
		return err
	}
	err = transaction.Commit()
	if err != nil {
		return err
	}

	return nil
}

func mappingUser(user *user.User, request updateUserRequest.Request) {
	user.UpdatedAt = sql.NullTime{Time: time.Now(), Valid: true}
	if request.Name != "" {
		user.Name = request.Name
	}
	if request.Email != "" {
		user.Email = sql.NullString{String: request.Email, Valid: true}
	}
	if request.RoleId != 0 {
		user.RoleId = request.RoleId
	}
	if request.Phone != "" {
		user.Phone = sql.NullString{String: request.Phone, Valid: true}
	}
	if request.Password != "" {
		user.Password = sql.NullString{String: request.Password, Valid: true}
	}
	if request.HoReCaId != 0 {
		user.HoReCaId = sql.NullInt16{Int16: int16(request.HoReCaId), Valid: true}
	}
	if request.Url != "" {
		user.PasswordRecoveryUrl = sql.NullString{String: request.Url, Valid: true}
	}
}
