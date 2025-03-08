package entity

import (
	"github.com/form3tech-oss/jwt-go"
	"time"
)

type AccessToken struct {
	User    User
	Token   *jwt.Token
	Expires time.Time
}
