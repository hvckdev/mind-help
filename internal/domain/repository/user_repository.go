package repository

import (
	"github.com/go-pg/pg/v10/orm"
	"mind-help/internal/domain/entity"
)

type UserRepository interface {
	GetById(id int) (entity.User, error)
	GetByEmail(email string) (entity.User, error)
	Save(newUser *entity.User) (orm.Model, error)
}
