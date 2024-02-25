package repository

import "mind-help/internal/domain/entities"

type UserRepository interface {
	GetById(id int) entities.User
	Save(newUser *entities.User) entities.User
}
