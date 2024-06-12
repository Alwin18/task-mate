package models

import "time"

type Attendance struct {
	ID         int64     `db:"id"`
	EmployeeID int64     `db:"employee_id"`
	Date       time.Time `db:"date"`
	CheckIn    time.Time `db:"check_in"`
	CheckOut   time.Time `db:"check_out"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

func (a *Attendance) TableName() string {
	return "attendance"
}
