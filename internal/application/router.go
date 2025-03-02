package application

import (
	"github.com/gin-gonic/gin"
	"mind-help/internal/infrastructure/controller"
	"mind-help/internal/infrastructure/repository"
	"mind-help/internal/infrastructure/service"
)

func (app *Application) RegisterRoutes(r *gin.Engine) {
	userController := controller.UserController{
		Service: service.UserService{
			Repository: &repository.UserRepository{
				Db: app.db,
			},
		},
	}

	r.POST("/api/v1/register", func(context *gin.Context) {
		userController.Register(context)
	})
}
