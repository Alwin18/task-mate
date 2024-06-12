package models

import "time"

type Employee struct {
	ID          int64     `db:"id"`
	UserID      int64     `db:"user_id"`
	FullName    string    `db:"fullname"`
	Address     string    `db:"address"`
	PhoneNumber string    `db:"phone_number"`
	NIP         string    `db:"nip"`
	Photo       string    `db:"photo"`
	Birthday    time.Time `db:"birthday"`
	JoinDate    time.Time `db:"join_date"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	Email       string    `db:"email"`
}

func (e *Employee) TableName() string {
	return "employees"
}
