package blog

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/model"
	"STUOJ/utils"
	"errors"
	"log"
)

// 提交记录统计
func GetStatistics(p model.Period) (model.BlogStatistics, error) {
	var err error
	var cbds []model.CountByDate
	var stats model.BlogStatistics

	// 检查时间范围
	err = p.Check()
	if err != nil {
		return model.BlogStatistics{}, err
	}

	// 统计博客数量
	stats.BlogCount, err = dao.CountBlogs()
	if err != nil {
		log.Println(err)
		return model.BlogStatistics{}, errors.New("统计博客数量失败")
	}

	// 统计博客数量
	cbds, err = dao.CountBlogsBetweenCreateTime(p.StartTime, p.EndTime)
	if err != nil {
		log.Println(err)
		return model.BlogStatistics{}, errors.New("统计博客数量失败")
	}
	stats.BlogCountByDate.FromCountByDate(cbds)
	utils.MapCountFillZero(&stats.BlogCountByDate, p.StartTime, p.EndTime)

	return stats, nil
}
