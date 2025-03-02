package repository

import "mind-help/internal/domain/entity"

type UserRepository interface {
	GetById(id int) entity.User
	Save(newUser *entity.User) entity.User
}
