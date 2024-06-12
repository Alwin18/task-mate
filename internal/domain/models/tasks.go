package models

import "time"

type Task struct {
	ID         int64     `db:"id"`
	ProjectID  int64     `db:"project_id"`
	EmployeeID int64     `db:"employee_id"`
	Name       string    `db:"name"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

func (t *Task) TableName() string {
	return "tasks"
}
