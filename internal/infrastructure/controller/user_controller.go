package controller

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/gin-gonic/gin"
	"mind-help/internal/infrastructure/request/user"
	"mind-help/internal/infrastructure/service"
	"net/http"
)

type UserController struct {
	Service service.UserService
}

func (userController *UserController) Register(c *gin.Context) {
	var body user.RegisterRequest

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createUser, err := userController.Service.CreateUser(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"result": createUser})
}

func (userController *UserController) Login(c *gin.Context) {
	var body user.LoginRequest

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jwt.New(jwt.SigningMethodHS256)
}
