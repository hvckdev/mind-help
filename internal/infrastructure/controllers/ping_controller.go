package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type PingController struct {
}

func (pingController *PingController) Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
