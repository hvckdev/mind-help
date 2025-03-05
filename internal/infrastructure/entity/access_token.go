package entity

import (
	"github.com/form3tech-oss/jwt-go"
	"mind-help/internal/domain/entity"
)

type AccessToken struct {
	User  entity.User
	Token jwt.Token
}
