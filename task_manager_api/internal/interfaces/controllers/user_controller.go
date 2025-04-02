package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zaahidali/task_manager_api/internal/domain/models"
	"github.com/zaahidali/task_manager_api/internal/interfaces/middleware"
	"github.com/zaahidali/task_manager_api/internal/usecases"
)

type UserController struct {
	usecase *usecases.UserUsecase
}

func NewUserController(u *usecases.UserUsecase) *UserController {
	return &UserController{usecase: u}
}

func (ctl *UserController) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := ctl.usecase.Register(context.Background(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User registered"})
}

func (ctl *UserController) Login(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := ctl.usecase.Login(context.Background(), input.Username, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	token, _ := middleware.GenerateToken(user.ID.Hex(), user.Username, user.Role)
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (ctl *UserController) Promote(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := ctl.usecase.Promote(context.Background(), input.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User promoted to admin"})
}
