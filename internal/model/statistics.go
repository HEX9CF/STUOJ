package model

import "time"

type CountByDate struct {
	Date  time.Time `json:"date"`
	Count uint64    `json:"count"`
}

type MapCountByDate map[string]uint64

type UserStatistics struct {
	RegisterCount MapCountByDate `json:"register_count"`
}
