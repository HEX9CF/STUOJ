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
	var err error
	var cbds []model.CountByDate
	var cbjss []model.CountByJudgeStatus
	var stats model.RecordStatistics

	// 检查时间范围
	if startTime.After(endTime) {
		return model.RecordStatistics{}, errors.New("开始时间不能晚于结束时间")
	}

	// 统计提交记录数量
	stats.SubmissionCount, err = dao.CountSubmissions()
	if err != nil {
		log.Println(err)
		return model.RecordStatistics{}, errors.New("统计提交记录数量失败")
	}

	// 统计评测点结果数量
	stats.JudgementCount, err = dao.CountJudgements()
	if err != nil {
		log.Println(err)
		return model.RecordStatistics{}, errors.New("统计评测点结果数量失败")
	}

	// 按评测状态统计提交记录数量
	cbjss, err = dao.CountSubmissionsGroupByStatus()
	if err != nil {
		log.Println(err)
		return model.RecordStatistics{}, errors.New("统计提交记录失败")
	}
	stats.SubmissionCountByStatus.FromCountByJudgeStatus(cbjss)

	// 按评测状态统计评测结果数量
	cbjss, err = dao.CountJudgementsGroupByStatus()
	if err != nil {
		log.Println(err)
		return model.RecordStatistics{}, errors.New("统计评测结果失败")
	}
	stats.JudgementCountByStatus.FromCountByJudgeStatus(cbjss)

	// 按日期统计提交记录数量
	cbds, err = dao.CountSubmissionsBetweenCreateTime(startTime, endTime)
	if err != nil {
		log.Println(err)
		return model.RecordStatistics{}, errors.New("统计提交记录失败")
	}
	stats.SubmissionCountByDate.FromCountByDate(cbds)

	return stats, nil
}
