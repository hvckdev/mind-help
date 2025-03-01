package application

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"log"
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

	log.Printf("%s:%s@%s:%s/%s?sslmode=disable", conf.User, conf.Password, conf.Host, conf.Port, conf.Database)
	m, err := migrate.New("file://migrations", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", conf.User, conf.Password, conf.Host, conf.Port, conf.Database))
	if err != nil {
		log.Panicf("Migrations init error: %s", err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Panicf("Up err: %v", err)
	}

	log.Println("Migrations success")
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
