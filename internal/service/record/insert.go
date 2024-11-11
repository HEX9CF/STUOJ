package record

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/model"
	"time"
)

// 插入提交记录
func InsertSubmission(s model.Submission) (uint64, error) {
	var err error

	updateTime := time.Now()
	s.UpdateTime = updateTime
	s.CreateTime = updateTime

	s.Id, err = dao.InsertSubmission(s)
	if err != nil {
		return 0, err
	}

	return s.Id, nil
}

// 插入评测结果
func InsertJudgement(j model.Judgement) (uint64, error) {
	var err error

	j.Id, err = dao.InsertJudgement(j)
	if err != nil {
		return 0, err
	}

	return j.Id, nil
}
