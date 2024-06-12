package repositories

import "task-mate/internal/domain/models"

type UserRepository interface {
	GetByID(id int) (*models.User, error)
}
