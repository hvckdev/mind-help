package service

import (
	"errors"
	"github.com/go-pg/pg/v10/orm"
	"github.com/golang-jwt/jwt/v5"
	"mind-help/internal/domain/entity"
	"mind-help/internal/domain/repository"
	"mind-help/internal/infrastructure/helpers"
	"mind-help/internal/infrastructure/request/user"
	"time"
)

type UserService struct {
	Repository repository.UserRepository
}

func (s *UserService) CreateUser(request user.RegisterRequest) (orm.Model, error) {
	hashedPassword, err := helpers.HashPassword(request.Password)

	if err != nil {
		return nil, err
	}

	model := &entity.User{
		Name:         request.Name,
		Email:        request.Email,
		Password:     hashedPassword,
		RegisteredAt: time.Now(),
		UpdatedAt:    time.Now(),
	}

	save, err := s.Repository.Save(model)
	if err != nil {
		return nil, err
	}

	return save, nil
}

func (s *UserService) Login(request user.LoginRequest) error {
	user, err := s.Repository.GetByEmail(request.Password)
	if err != nil {
		return err
	}

	hashedPassword, err := helpers.HashPassword(request.Password)
	if hashedPassword != user.Password {
		return errors.New("")
	}

	token := jwt.New(jwt.SigningMethodHS256)

	return nil
}
