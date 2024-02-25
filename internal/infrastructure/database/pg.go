package database

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"mind-help/internal/application"
)

type Database struct {
	db *pg.DB
}

func (database Database) Connect() {
	config := application.Config{}.Database()

	database.db = pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Host, config.Port),
		User:     config.User,
		Password: config.Password,
	})

	defer func(db *pg.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(database.db)
}
