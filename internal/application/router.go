package application

import (
	"github.com/gin-gonic/gin"
	"mind-help/internal/infrastructure/controllers"
)

func (app *Application) RegisterRoutes(r *gin.Engine) {
	pingController := controllers.PingController{}

	r.GET("/ping", func(context *gin.Context) {
		pingController.Index(context)
	})
}
