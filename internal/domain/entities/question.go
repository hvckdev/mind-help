package entities

import "time"

type Question struct {
	ID        int64     `json:"id,omitempty"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Owner     User      `json:"owner"`
}
