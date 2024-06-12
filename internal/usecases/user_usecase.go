package usecases

import (
	"task-mate/internal/domain/models"
	"task-mate/internal/domain/repositories"
)

type UserUseCase struct {
	UserRepository repositories.UserRepository
}

func (uc *UserUseCase) GetUserByID(id int) (user *models.User, err error) {
	return
}
