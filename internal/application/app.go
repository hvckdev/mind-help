package application

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"mind-help/internal/infrastructure/database"
)

type Application struct {
}

func (app *Application) Start() {
	app.loadEnvironmentVariables()
	app.connectToDB()

	app.runRouter()
}

func (app *Application) loadEnvironmentVariables() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func (app *Application) connectToDB() {
	database.Database{}.Connect()
}

func (app *Application) runRouter() {
	r := app.createRouter()
	app.RegisterRoutes(r)
	err := r.Run()

	if err != nil {
		panic(err)
	}
}

func (app *Application) createRouter() *gin.Engine {
	r := gin.Default()

	return r
}
