package user

import "database/sql"

type User struct {
	UserId              int            `db:"user_id"`
	Name                string         `db:"name"`
	RoleId              int            `db:"role_id"`
	Phone               sql.NullString `db:"phone"`
	Password            sql.NullString `db:"password"`
	Email               sql.NullString `db:"email"`
	HoReCaId            sql.NullInt16  `db:"horeca_id"`
	PasswordRecoveryUrl sql.NullString `db:"password_recovery_url"`
	Messenger           sql.NullString `db:"messenger"`
	CreatedAt           sql.NullTime   `db:"created_at"`
	UpdatedAt           sql.NullTime   `db:"updated_at"`
}
