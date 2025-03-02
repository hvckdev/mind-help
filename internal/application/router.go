package application

import (
	"github.com/gin-gonic/gin"
	"mind-help/internal/infrastructure/controller"
)

func (app *Application) RegisterRoutes(r *gin.Engine) {
	pingController := controller.PingController{
		DB: app.db,
	}

	r.GET("/ping", func(context *gin.Context) {
		pingController.Index(context)
	})
}
