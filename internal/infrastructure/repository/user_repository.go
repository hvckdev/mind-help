package repository

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"mind-help/internal/domain/entity"
)

type UserRepository struct {
	DB *pg.DB
}

func (r *UserRepository) GetById(userId int) (entity.User, error) {
	var user entity.User

	err := r.DB.Model(&user).Where("id = ?", userId).Select()
	if err != nil {
		return entity.User{}, err
	}

	return user, err
}

func (r *UserRepository) GetByEmail(email string) (entity.User, error) {
	var user entity.User

	err := r.DB.Model(&user).Where("id = ?", email).Select()
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (r *UserRepository) Save(newUser *entity.User) (orm.Model, error) {
	insert, err := r.DB.Model(newUser).Insert()
	if err != nil {
		return nil, err
	}

	return insert.Model(), nil
}
