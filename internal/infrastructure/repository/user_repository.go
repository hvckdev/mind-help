package repository

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"mind-help/internal/domain/entity"
)

type UserRepository struct {
	DB *pg.DB
}

func (r *UserRepository) GetById(id int) entity.User {
	//TODO implement me
	panic("implement me")
}

func (r *UserRepository) Save(newUser *entity.User) (orm.Model, error) {
	insert, err := r.DB.Model(newUser).Insert()
	if err != nil {
		return nil, err
	}

	return insert.Model(), nil
}
