package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"mind-help/internal/infrastructure/request/user"
	"net/http"
)

type UserController struct {
	DB *pg.DB
}

func (userController UserController) register(c *gin.Context) {
	var body user.RegisterRequest

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

}
