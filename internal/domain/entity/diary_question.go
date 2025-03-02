package entity

type DiaryQuestion struct {
	Question Question `json:"question"`
	Diary    Diary    `json:"diary"`
}
