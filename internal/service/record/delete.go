package record

import (
	"STUOJ/internal/dao"
)

// 根据ID删除提交记录
func DeleteSubmissionById(id uint64) error {
	err := dao.DeleteSubmissionById(id)
	if err != nil {
		return err
	}

	return nil
}

// 根据提交ID查询评测结果
func DeleteJudgementBySubmissionId(id uint64) error {
	err := dao.DeleteJudgementBySubmissionId(id)
	if err != nil {
		return err
	}

	return nil
}
