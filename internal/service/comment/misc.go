package comment

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/model"
	"errors"
	"log"
	"time"
)

// 提交记录统计
func GetStatistics(startTime time.Time, endTime time.Time) (model.CommentStatistics, error) {
	var err error
	var cbds []model.CountByDate
	var stats model.CommentStatistics

	// 检查时间范围
	if startTime.After(endTime) {
		return model.CommentStatistics{}, errors.New("开始时间不能晚于结束时间")
	}

	// 统计评论数量
	stats.CommentCount, err = dao.CountBlogs()
	if err != nil {
		log.Println(err)
		return model.CommentStatistics{}, errors.New("统计评论数量失败")
	}

	// 统计评论数量
	cbds, err = dao.CountCommentsBetweenCreateTime(startTime, endTime)
	if err != nil {
		log.Println(err)
		return model.CommentStatistics{}, errors.New("统计评论数量失败")
	}
	stats.CommentCountByDate.FromCountByDate(cbds)
	stats.CommentCountByDate.FillZero(startTime, endTime)

	return stats, nil
}
