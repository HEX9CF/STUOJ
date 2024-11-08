package dao

import (
	"STUOJ/internal/db"
	"STUOJ/internal/model"
	"time"
)

// 插入提交记录
func InsertSubmission(s model.Submission) (uint64, error) {
	updateTime := time.Now()
	s.UpdateTime = updateTime
	s.CreateTime = updateTime
	tx := db.Db.Create(&s)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return s.Id, nil
}

// 查询所有提交记录（不返回源代码）
func SelectAllSubmissions() ([]model.Submission, error) {
	var submissions []model.Submission

	tx := db.Db.Omit("source_code").Find(&submissions)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return submissions, nil
}

// 根据ID查询提交记录
func SelectSubmissionById(id uint64) (model.Submission, error) {
	var s model.Submission

	tx := db.Db.Where("id = ?", id).First(&s)
	if tx.Error != nil {
		return model.Submission{}, tx.Error
	}

	return s, nil
}

// 根据用户ID查询提交记录（不返回源代码）
func SelectSubmissionsByUserId(userId uint64) ([]model.Submission, error) {
	var submissions []model.Submission

	tx := db.Db.Omit("source_code").Where("user_id = ?", userId).Find(&submissions)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return submissions, nil
}

// 根据题目ID查询提交记录（不返回源代码）
func SelectSubmissionsByProblemId(problemId uint64) ([]model.Submission, error) {
	var submissions []model.Submission

	tx := db.Db.Omit("source_code").Where("problem_id = ?", problemId).Find(&submissions)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return submissions, nil
}

// 更新提交记录
func UpdateSubmissionById(s model.Submission) error {
	updateTime := time.Now()
	tx := db.Db.Model(&s).Where("id = ?", s.Id).Updates(map[string]interface{}{
		"user_id":     s.UserId,
		"problem_id":  s.ProblemId,
		"status":      s.Status,
		"score":       s.Score,
		"language_id": s.LanguageId,
		"length":      s.Length,
		"memory":      s.Memory,
		"time":        s.Time,
		"source_code": s.SourceCode,
		"update_time": updateTime,
	})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// 根据ID更新提交记录状态更新时间
func UpdateSubmissionUpdateTimeById(id uint64) error {
	updateTime := time.Now()
	tx := db.Db.Model(&model.Submission{}).Where("id = ?", id).Updates(map[string]interface{}{
		"update_time": updateTime,
	})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// 根据ID删除提交记录
func DeleteSubmissionById(id uint64) error {
	tx := db.Db.Where("id = ?", id).Delete(&model.Submission{})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
