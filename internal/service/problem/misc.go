package problem

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/model"
	"errors"
	"log"
	"time"
)

// 提交记录统计
func GetStatistics(startTime time.Time, endTime time.Time) (model.ProblemStatistics, error) {
	var stats model.ProblemStatistics

	// 检查时间范围
	if startTime.After(endTime) {
		return model.ProblemStatistics{}, errors.New("开始时间不能晚于结束时间")
	}

	// 统计用户注册数量
	countByCreateTime, err := dao.CountProblemsBetweenCreateTime(startTime, endTime)
	if err != nil {
		log.Println(err)
		return model.ProblemStatistics{}, errors.New("统计题目数量失败")
	}

	stats.AddCountByDate = make(model.MapCountByDate)
	for _, v := range countByCreateTime {
		date := v.Date.Format(model.LayoutCountByDate)
		stats.AddCountByDate[date] = v.Count
	}

	return stats, nil
}
