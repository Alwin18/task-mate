package models

import "time"

type User struct {
	ID        int64     `db:"id"`
	Username  string    `db:"username"`
	Password  string    `db:"password"`
	RoleID    int64     `db:"role_id"`
	IsActive  bool      `db:"is_active"`
	UpdatedAt time.Time `db:"updated_at"`
	CreatedAt time.Time `db:"created_at"`
}

func (u *User) TableName() string {
	return "users"
}
