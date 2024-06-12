package models

import "time"

type Role struct {
	ID        int64     `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Flag      string    `db:"flag"`
	Role      string    `db:"role"`
}

func (r *Role) TableName() string {
	return "roles"
}

func (r *Role) GetID() int64 {
	return r.ID
}
