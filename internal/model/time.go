package model

import (
	"errors"
	"log"
	"time"
)

// 时间范围
type Period struct {
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

// 检查时间范围
func (p Period) Check() error {
	if p.StartTime.After(p.EndTime) {
		return errors.New("开始时间不能晚于结束时间")
	}
	return nil
}

// 从字符串解析时间范围
func (p *Period) FromString(startTimeStr string, endTimeStr string, layout string) error {
	var err error

	if startTimeStr == "" || endTimeStr == "" {
		return errors.New("参数错误")
	}

	p.StartTime, err = time.Parse(layout, startTimeStr)
	if err != nil {
		log.Println(err)
		return errors.New("开始时间格式错误")
	}
	p.EndTime, err = time.Parse(layout, endTimeStr)
	if err != nil {
		log.Println(err)
		return errors.New("结束时间格式错误")
	}

	return nil
}
