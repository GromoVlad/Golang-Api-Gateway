package user

import (
	"database/sql"
)

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	UserId              int            `db:"user_id" json:"user_id" example:"42" format:"int"`
	Name                string         `db:"name" json:"name" example:"John Doe" format:"string"`
	RoleId              int            `db:"role_id" json:"role_id" example:"42" format:"int"`
	Phone               sql.NullString `db:"phone" json:"phone" example:"89998885566" format:"string" swaggertype:"string"`
	Password            sql.NullString `db:"password" json:"password" example:"*************" format:"string" swaggertype:"string"`
	Email               sql.NullString `db:"email" json:"email" example:"example@example.com" format:"string" swaggertype:"string"`
	HoReCaId            sql.NullInt16  `db:"horeca_id" json:"horeca_id" example:"42" format:"int" swaggertype:"integer"`
	PasswordRecoveryUrl sql.NullString `db:"password_recovery_url" json:"password_recovery_url" example:"http://localhost:8080/password_recovery_url" format:"string" swaggertype:"string"`
	Messenger           sql.NullString `db:"messenger" json:"messenger" example:"telegram" format:"string" swaggertype:"string"`
	CreatedAt           sql.NullTime   `db:"created_at" json:"created_at" example:"2022-01-01 00:00:00" format:"string" swaggertype:"string"`
	UpdatedAt           sql.NullTime   `db:"updated_at" json:"updated_at" example:"2022-01-01 00:00:00" format:"string" swaggertype:"string"`
}
