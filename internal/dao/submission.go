package dao

import (
	"STUOJ/internal/db"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"time"
)

// 插入提交记录
func InsertSubmission(s entity.Submission) (uint64, error) {
	updateTime := time.Now()
	s.UpdateTime = updateTime
	s.CreateTime = updateTime
	tx := db.Db.Create(&s)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return s.Id, nil
}

// 查询所有提交记录
func SelectAllSubmissions() ([]entity.Submission, error) {
	var submissions []entity.Submission

	tx := db.Db.Find(&submissions)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return submissions, nil
}

// 根据ID查询提交记录
func SelectSubmissionById(id uint64) (entity.Submission, error) {
	var s entity.Submission

	tx := db.Db.Where("id = ?", id).First(&s)
	if tx.Error != nil {
		return entity.Submission{}, tx.Error
	}

	return s, nil
}

// 根据用户ID查询提交记录（不返回源代码）
func SelectSubmissionsByUserId(userId uint64) ([]entity.Submission, error) {
	var submissions []entity.Submission

	tx := db.Db.Where("user_id = ?", userId).Find(&submissions)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return submissions, nil
}

// 根据题目ID查询提交记录（不返回源代码）
func SelectSubmissionsByProblemId(problemId uint64) ([]entity.Submission, error) {
	var submissions []entity.Submission

	tx := db.Db.Where("problem_id = ?", problemId).Find(&submissions)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return submissions, nil
}

// 更新提交记录
func UpdateSubmissionById(s entity.Submission) error {
	tx := db.Db.Model(&s).Where("id = ?", s.Id).Updates(s)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// 根据ID更新提交记录的更新时间
func UpdateSubmissionUpdateTimeById(id uint64) error {
	tx := db.Db.Model(&entity.Submission{}).Where("id = ?", id).Update("update_time", time.Now())
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// 根据ID删除提交记录
func DeleteSubmissionById(id uint64) error {
	tx := db.Db.Where("id = ?", id).Delete(&entity.Submission{})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// 统计提交信息数量
func CountSubmissions() (int64, error) {
	var count int64

	tx := db.Db.Model(&entity.Submission{}).Count(&count)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return count, nil
}

// 根据创建时间统计用户数量
func CountSubmissionsBetweenCreateTime(startTime time.Time, endTime time.Time) ([]model.CountByDate, error) {
	var countByDate []model.CountByDate

	tx := db.Db.Model(&entity.Submission{}).Where("create_time between ? and ?", startTime, endTime).Select("date(create_time) as date, count(*) as count").Group("date").Scan(&countByDate)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return countByDate, nil
}
