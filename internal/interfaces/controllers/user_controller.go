package controllers

import (
	"net/http"
	"strconv"
	"task-mate/internal/usecases"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUseCase *usecases.UserUseCase
}

func (uc *UserController) GetUserByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := uc.UserUseCase.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
