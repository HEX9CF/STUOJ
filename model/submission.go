package model

import "time"

type Submission struct {
	Id         uint64    `json:"id"`
	UserID     uint64    `json:"user_id"`
	QuestionID uint64    `json:"question_id"`
	Status     uint64    `json:"status"`
	Score      uint64    `json:"score"`
	SubmitTime time.Time `json:"submit_time"`
	LanguageID uint64    `json:"language_id"`
	Length     uint64    `json:"length"`
	Memory     uint64    `json:"memory"`
	Time       uint64    `json:"time"`
}
