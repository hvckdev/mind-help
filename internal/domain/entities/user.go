package entities

import "time"

type User struct {
	ID           int64     `json:"id,omitempty"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	RegisteredAt time.Time `json:"registeredAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type Question struct {
	ID        int64     `json:"id,omitempty"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Owner     User      `json:"owner"`
}

type Diary struct {
	ID        int64     `json:"ID"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Owner     User      `json:"owner"`
}

type DiaryQuestion struct {
	Question Question `json:"question"`
	Diary    Diary    `json:"diary"`
}
