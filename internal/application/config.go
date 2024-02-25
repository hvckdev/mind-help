package application

import "os"

type Config struct {
}

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
}

func (config Config) Database() *Database {
	return &Database{
		Host:     os.Getenv("DATABASE_HOST"),
		Port:     os.Getenv("DATABASE_PORT"),
		User:     os.Getenv("DATABASE_USER"),
		Password: os.Getenv("DATABASE_PASSWORD"),
	}
}
