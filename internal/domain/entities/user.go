package entities

import "time"

type User struct {
	ID           int64     `json:"id,omitempty"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	RegisteredAt time.Time `json:"registeredAt,pg:registered_at"`
	UpdatedAt    time.Time `json:"updatedAt,pg:updated_at"`
}
