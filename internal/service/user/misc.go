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
	var err error
	var stats model.UserStatistics

	// 检查时间范围
	if startTime.After(endTime) {
		return model.UserStatistics{}, errors.New("开始时间不能晚于结束时间")
	}

	// 统计用户数量
	stats.UserCount, err = dao.CountUsers()
	if err != nil {
		log.Println(err)
		return model.UserStatistics{}, errors.New("统计用户数量失败")
	}

	// 统计用户注册数量
	cbds, err := dao.CountUsersBetweenCreateTime(startTime, endTime)
	if err != nil {
		log.Println(err)
		return model.UserStatistics{}, errors.New("统计用户注册数量失败")
	}
	stats.RegisterCountByDate.FromStruct(cbds)

	return stats, nil
}
