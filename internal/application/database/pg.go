package database

import (
	"github.com/go-pg/pg/v10"
)

type Database struct {
	db *pg.DB
}

func (database Database) Connect(opts *pg.Options) *pg.DB {
	db := pg.Connect(opts)

	return db
}

func (database Database) GetConnection() *pg.DB {
	return database.db
}
