package repository

import (
	"errors"
	"task-mate/internal/domain/models"
)

type InMemoryUserRepository struct {
	users map[int]*models.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[int]*models.User),
	}
}

func (repo *InMemoryUserRepository) GetByID(id int) (user *models.User, err error) {
	if user, exist := repo.users[id]; exist {
		return user, nil
	}
	return nil, errors.New("user not found")
}
