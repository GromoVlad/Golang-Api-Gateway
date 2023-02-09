package userRepository

import (
	"database/sql"
	"errors"
	"fmt"
	"gin_tonic/internal/database/DB"
	"gin_tonic/internal/enums/messenger"
	"gin_tonic/internal/models/user"
	"gin_tonic/internal/requests/user/createUserRequest"
	"gin_tonic/internal/requests/user/listUserRequest"
	"gin_tonic/internal/requests/user/updateUserRequest"
	"gin_tonic/internal/support/localContext"
	_ "github.com/lib/pq"
	"time"
)

func FindUser(context localContext.LocalContext, userId int) user.User {
	connect := DB.Connect()
	defer connect.Close()

	var findUser user.User
	_ = connect.Get(&findUser, "SELECT * FROM users.users WHERE user_id = $1", userId)
	if findUser.UserId == 0 {
		context.NotFoundError(
			errors.New(fmt.Sprintf("Пользователь с идентификатором %d не зарегистрирован в системе", userId)),
		)
	}
	return findUser
}

func FindUsers(context localContext.LocalContext, request listUserRequest.Request) ([]user.User, int) {
	connect := DB.Connect()
	defer connect.Close()

	var users []user.User
	var err, errTotal error
	var total int

	if request.Search != "" {
		query := "SELECT * FROM users.users WHERE name ilike $1 LIMIT $2 OFFSET $3"
		err = connect.Select(&users, query, "%"+request.Search+"%", request.Limit, request.Offset)
		errTotal = connect.
			QueryRow("SELECT COUNT(user_id) AS total FROM users.users WHERE name ilike $1", "%"+request.Search+"%").
			Scan(&total)
	} else {
		query := "SELECT * FROM users.users LIMIT $1 OFFSET $2"
		err = connect.Select(&users, query, request.Limit, request.Offset)
		errTotal = connect.QueryRow("SELECT COUNT(user_id) AS total FROM users.users").Scan(&total)
	}

	context.InternalServerError(err)
	context.InternalServerError(errTotal)

	totalPage := calcTotalPage(request.Limit, total)

	return users, totalPage
}

func CreateUser(context localContext.LocalContext, dto createUserRequest.DTO) {
	connect := DB.Connect()
	defer connect.Close()

	var findUser user.User
	if dto.Email != "" {
		_ = connect.Get(&findUser, "SELECT user_id FROM users.users WHERE email = $1", dto.Email)
		if findUser.UserId != 0 {
			context.AlreadyExistsError(errors.New("Пользователь с email " + dto.Email + " уже зарегистрирован в системе"))
		}
	}

	transaction := connect.MustBegin()
	_, err := transaction.NamedExec(
		"INSERT INTO users.users (name, role_id, phone, password, email, venue_id, password_recovery_url, messenger, created_at, updated_at) "+
			"VALUES (:name, :role_id, :phone, :password, :email, :venue_id, :password_recovery_url, :messenger, :created_at, :updated_at)",
		&user.User{
			Name:                dto.Name,
			RoleId:              dto.RoleId,
			Phone:               sql.NullString{String: dto.Phone, Valid: dto.Phone != ""},
			Password:            dto.Password,
			Email:               dto.Email,
			VenueId:             sql.NullInt16{Int16: int16(dto.VenueId), Valid: dto.VenueId != 0},
			PasswordRecoveryUrl: sql.NullString{},
			Messenger:           sql.NullString{String: messenger.TELEGRAM, Valid: true},
			CreatedAt:           sql.NullTime{Time: time.Now(), Valid: true},
			UpdatedAt:           sql.NullTime{},
		},
	)
	context.StatusConflictError(err)

	err = transaction.Commit()
	context.InternalServerError(err)
}

func UpdateUser(context localContext.LocalContext, request updateUserRequest.Request) {
	connect := DB.Connect()
	defer connect.Close()

	findUser := FindUser(context, request.UserId)

	mappingUser(&findUser, request)

	transaction := connect.MustBegin()
	_, err := transaction.NamedExec(
		"UPDATE users.users SET updated_at = :updated_at, name = :name, role_id = :role_id, "+
			"phone = :phone, password = :password, email = :email, venue_id = :venue_id, "+
			"password_recovery_url = :password_recovery_url WHERE user_id = :user_id",
		&findUser,
	)
	context.StatusConflictError(err)

	err = transaction.Commit()
	context.InternalServerError(err)
}

func DeleteUser(context localContext.LocalContext, userId int) {
	connect := DB.Connect()
	defer connect.Close()

	FindUser(context, userId)

	transaction := connect.MustBegin()
	_, err := transaction.NamedExec("DELETE FROM users.users WHERE user_id = :user_id", &user.User{UserId: userId})
	context.StatusConflictError(err)
	err = transaction.Commit()
	context.InternalServerError(err)
}

func FindOrFailByEmail(context localContext.LocalContext, email string) user.User {
	connect := DB.Connect()
	defer connect.Close()

	var findUser user.User
	if email != "" {
		_ = connect.Get(&findUser, "SELECT * FROM users.users WHERE email = $1", email)
		if findUser.UserId == 0 {
			context.NotFoundError(
				errors.New(fmt.Sprintf("Пользователь с email %s не зарегистрирован в системе", email)),
			)
		}
	}
	return findUser
}

func mappingUser(user *user.User, request updateUserRequest.Request) {
	user.UpdatedAt = sql.NullTime{Time: time.Now(), Valid: true}
	if request.Name != "" {
		user.Name = request.Name
	}
	if request.Email != "" {
		user.Email = request.Email
	}
	if request.Password != "" {
		user.Password = request.Password
	}
	if request.RoleId != 0 {
		user.RoleId = request.RoleId
	}
	if request.Phone != "" {
		user.Phone = sql.NullString{String: request.Phone, Valid: true}
	}
	if request.VenueId != 0 {
		user.VenueId = sql.NullInt16{Int16: int16(request.VenueId), Valid: true}
	}
	if request.Url != "" {
		user.PasswordRecoveryUrl = sql.NullString{String: request.Url, Valid: true}
	}
}

func calcTotalPage(limit int, total int) int {
	var count, countRemainderOfDivision int
	count = total / limit
	countRemainderOfDivision = total % limit
	if countRemainderOfDivision > 0 {
		return count + 1
	} else {
		return count
	}
}
