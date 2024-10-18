package model

import "time"

type Problem struct {
	Id         uint64    `json:"id"`
	Title      string    `json:"title"`
	Source     string    `json:"source"`
	Difficulty uint64    `json:"difficulty"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}
