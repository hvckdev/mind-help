package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"mind-help/internal/domain/entity"
	"net/http"
	"time"
)

type PingController struct {
	DB *pg.DB
}

func (pingController *PingController) Index(c *gin.Context) {
	model := &entity.User{
		ID:           0,
		Name:         "sdfsdf",
		Email:        "sdfsdfsdf",
		Password:     "sdfsdfsdf",
		RegisteredAt: time.Now(),
		UpdatedAt:    time.Now(),
	}

	_, err := pingController.DB.Model(model).Insert()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
