package user

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/model"
	"errors"
	"log"
	"time"
)

// 统计用户
func GetStatistics(startTime time.Time, endTime time.Time) (model.UserStatistics, error) {
	var stats model.UserStatistics

	// 统计用户注册数量
	countByCreateTime, err := dao.CountUsersBetweenCreateTime(startTime, endTime)
	if err != nil {
		log.Println(err)
		return model.UserStatistics{}, errors.New("统计用户注册数量失败")
	}

	stats.RegisterCount = make(map[string]uint64)
	for _, v := range countByCreateTime {
		date := v.Date.Format("2006-01-02")
		stats.RegisterCount[date] = v.Count
	}

	return stats, nil
}
