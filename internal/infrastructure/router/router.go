package router

import (
	"task-mate/internal/interfaces/controllers"
	"task-mate/internal/interfaces/repository"
	"task-mate/internal/usecases"

	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.Engine) {
	// Setup repository
	userRepository := repository.NewInMemoryUserRepository()

	// Setup use case
	userUseCase := &usecases.UserUseCase{UserRepository: userRepository}

	// Setup controller
	userController := &controllers.UserController{UserUseCase: userUseCase}

	r.GET("/users/:id", userController.GetUserByID)
}

func Route() *gin.Engine {
	r := gin.Default()
	UserRoute(r)
	return r
}
