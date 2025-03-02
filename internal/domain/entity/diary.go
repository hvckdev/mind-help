package entity

import "time"

type Diary struct {
	ID        int64     `json:"ID"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Owner     User      `json:"owner"`
}
