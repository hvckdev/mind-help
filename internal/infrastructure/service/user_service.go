package service

import (
	"github.com/go-pg/pg/v10/orm"
	"mind-help/internal/domain/entity"
	"mind-help/internal/domain/repository"
	"mind-help/internal/infrastructure/request/user"
	"time"
)

type UserService struct {
	Repository repository.UserRepository
}

func (s *UserService) CreateUser(request user.RegisterRequest) (orm.Model, error) {
	model := &entity.User{
		Name:         request.Name,
		Email:        request.Email,
		Password:     request.Password,
		RegisteredAt: time.Now(),
		UpdatedAt:    time.Now(),
	}

	save, err := s.Repository.Save(model)
	if err != nil {
		return nil, err
	}

	return save, nil
}
