package dao

import (
	"STUOJ/internal/db"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"time"

	"gorm.io/gorm"
)

type SubmissionWhere struct {
	ProblemId  model.Field[uint64]
	UserId     model.Field[uint64]
	LanguageId model.Field[uint64]
	StartTime  model.Field[time.Time]
	EndTime    model.Field[time.Time]
	Status     model.Field[uint64]
}

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
func SelectSubmissions(condition SubmissionWhere, page uint64, size uint64) ([]entity.Submission, error) {
	var submissions []entity.Submission

	where := generateSubmissionWhereCondition(condition)
	tx := db.Db.Offset(int((page - 1) * size)).Limit(int(size))
	tx = where(tx)
	tx = tx.Find(&submissions)
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
func CountSubmissions(condition SubmissionWhere) (uint64, error) {
	var count int64
	where := generateSubmissionWhereCondition(condition)
	tx := db.Db.Model(&entity.Submission{})
	tx = where(tx)
	tx = tx.Count(&count)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return uint64(count), nil
}

// 按评测状态统计提交信息数量
func CountSubmissionsGroupByStatus() ([]model.CountByJudgeStatus, error) {
	var counts []model.CountByJudgeStatus

	tx := db.Db.Model(&entity.Submission{}).Select("status, count(*) as count").Group("status").Scan(&counts)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return counts, nil
}

// 按语言ID统计提交信息数量
func CountSubmissionsGroupByLanguageId() ([]model.CountByLanguage, error) {
	var counts []model.CountByLanguage

	tx := db.Db.Model(&entity.Submission{}).Select("language_id, count(*) as count").Group("language_id").Scan(&counts)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return counts, nil
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

func generateSubmissionWhereCondition(con SubmissionWhere) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		whereClause := map[string]interface{}{}

		if con.ProblemId.Exist() {
			whereClause["problem_id"] = con.ProblemId.Value()
		}
		if con.UserId.Exist() {
			whereClause["user_id"] = con.UserId.Value()
		}
		if con.LanguageId.Exist() {
			whereClause["language_id"] = con.LanguageId.Value()
		}
		if con.Status.Exist() {
			whereClause["status"] = con.Status.Value()
		}
		where := db.Where(whereClause)
		if con.StartTime.Exist() {
			where.Where("create_time >= ?", con.StartTime.Value())
		}
		if con.EndTime.Exist() {
			where.Where("create_time <= ?", con.EndTime.Value())
		}
		return where
	}
}
