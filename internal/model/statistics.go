package model

import "time"

type CountByDate struct {
	Date  time.Time `json:"date"`
	Count uint      `json:"count"`
}
