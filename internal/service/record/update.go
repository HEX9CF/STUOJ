package record

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/model"
	"errors"
	"log"
	"time"
)

// 更新提交记录
func UpdateSubmissionById(s model.Submission) error {
	updateTime := time.Now()
	s.UpdateTime = updateTime

	err := dao.UpdateSubmissionById(s)
	if err != nil {
		return err
	}

	return nil
}

// 根据ID更新提交记录状态更新时间
func UpdateSubmissionUpdateTimeById(id uint64) error {
	updateTime := time.Now()

	s := model.Submission{
		Id:         id,
		UpdateTime: updateTime,
	}

	err := dao.UpdateSubmissionById(s)
	if err != nil {
		return err
	}

	return nil
}

// 更新评测结果
func UpdateJudgementById(j model.Judgement) error {
	err := dao.UpdateJudgementById(j)
	if err != nil {
		return err
	}

	// 更新提交记录状态更新时间
	err = UpdateSubmissionUpdateTimeById(j.SubmissionId)
	if err != nil {
		log.Println(err)
		return errors.New("更新提交记录状态更新时间失败")
	}

	return nil
}
