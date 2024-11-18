package dao

import (
	"STUOJ/internal/db"
	"STUOJ/internal/entity"
)

// 插入评测结果
func InsertJudgement(j entity.Judgement) (uint64, error) {
	tx := db.Db.Create(&j)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return j.Id, nil
}

// 根据提交ID查询评测结果
func SelectJudgementsBySubmissionId(sid uint64) ([]entity.Judgement, error) {
	var judgements []entity.Judgement

	tx := db.Db.Where("submission_id = ?", sid).Find(&judgements)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return judgements, nil
}

// 根据提交ID查询评测结果
func DeleteJudgementBySubmissionId(id uint64) error {
	tx := db.Db.Where("submission_id = ?", id).Delete(&entity.Judgement{})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// 更新评测结果
func UpdateJudgementById(j entity.Judgement) error {
	tx := db.Db.Model(&j).Where("id = ?", j.Id).Updates(j)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// 统计评测结果数量
func CountJudgements() (int64, error) {
	var count int64

	tx := db.Db.Model(&entity.Judgement{}).Count(&count)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return count, nil
}
