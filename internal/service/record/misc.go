package record

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/model"
	"errors"
	"log"
	"time"
)

// 提交记录统计
func GetStatistics(startTime time.Time, endTime time.Time) (model.RecordStatistics, error) {
	var stats model.RecordStatistics

	// 检查时间范围
	if startTime.After(endTime) {
		return model.RecordStatistics{}, errors.New("开始时间不能晚于结束时间")
	}

	// 统计用户注册数量
	countByCreateTime, err := dao.CountSubmissionsBetweenCreateTime(startTime, endTime)
	if err != nil {
		log.Println(err)
		return model.RecordStatistics{}, errors.New("统计提交记录失败")
	}

	stats.SubmitCount = make(map[string]uint64)
	for _, v := range countByCreateTime {
		date := v.Date.Format("2006-01-02")
		stats.SubmitCount[date] = v.Count
	}

	return stats, nil
}
