package application

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/joho/godotenv"
	"mind-help/internal/application/database"
)

type Application struct {
	db *pg.DB
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
	conf := Config{}.Database()

	app.db = database.Database{}.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%s", conf.Host, conf.Port),
		User:     conf.User,
		Password: conf.Password,
		Database: conf.Database,
	})
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
