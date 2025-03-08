package repository

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/go-pg/pg/v10"
	"mind-help/internal/domain/entity"
)

type TokenRepository struct {
	DB *pg.DB
}

func (r *TokenRepository) CreateForUser(u entity.User) (*jwt.Token, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	model := entity.AccessToken{
		User:  u,
		Token: token,
	}

	_, err := r.DB.Model(model).Insert()
	if err != nil {
		return nil, err
	}

	return token, nil
}
