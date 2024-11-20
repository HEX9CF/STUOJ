package blog

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/model"
	"errors"
	"log"
	"time"
)

// 提交记录统计
func GetStatistics(startTime time.Time, endTime time.Time) (model.BlogStatistics, error) {
	var err error
	var cbds []model.CountByDate
	var stats model.BlogStatistics

	// 检查时间范围
	if startTime.After(endTime) {
		return model.BlogStatistics{}, errors.New("开始时间不能晚于结束时间")
	}

	// 统计博客数量
	stats.BlogCount, err = dao.CountBlogs()
	if err != nil {
		log.Println(err)
		return model.BlogStatistics{}, errors.New("统计题目数量失败")
	}

	// 统计添加题目数量
	cbds, err = dao.CountBlogsBetweenCreateTime(startTime, endTime)
	if err != nil {
		log.Println(err)
		return model.BlogStatistics{}, errors.New("统计添加题目数量失败")
	}
	stats.BlogCountByDate.FromCountByDate(cbds)
	stats.BlogCountByDate.FillZero(startTime, endTime)

	return stats, nil
}
