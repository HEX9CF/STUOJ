package db

import (
	"STUOJ/model"
	"log"
	"time"
)

// 查询所有提交记录（不返回源代码）
func SelectAllSubmissions() ([]model.Submission, error) {
	sql := "SELECT id, user_id, problem_id, status, score, language_id, length, memory, time, create_time, update_time FROM tbl_submission"
	rows, err := Mysql.Query(sql)
	log.Println(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// 遍历查询结果
	submissions := make([]model.Submission, 0)
	for rows.Next() {
		var submission model.Submission
		var createTimeStr, updateTimeStr string

		err := rows.Scan(&submission.Id, &submission.UserId, &submission.ProblemId, &submission.Status, &submission.Score, &submission.LanguageId, &submission.Length, &submission.Memory, &submission.Time, &createTimeStr, &updateTimeStr)
		if err != nil {
			return nil, err
		}

		// 时间格式转换
		timeLayout := "2006-01-02 15:04:05"
		submission.CreateTime, err = time.Parse(timeLayout, createTimeStr)
		submission.UpdateTime, err = time.Parse(timeLayout, updateTimeStr)
		if err != nil {
			return nil, err
		}

		// 不返回源代码
		submission.SourceCode = ""

		//log.Println(submission)
		submissions = append(submissions, submission)
	}
	return submissions, nil
}

// 根据ID查询提交记录
func SelectSubmissionById(id uint64) (model.Submission, error) {
	var submission model.Submission
	var createTimeStr, updateTimeStr string

	submission.Id = id

	sql := "SELECT user_id, problem_id, status, score, language_id, length, memory, time, source_code, create_time, update_time FROM tbl_submission WHERE id = ? LIMIT 1"
	err := Mysql.QueryRow(sql, id).Scan(&submission.UserId, &submission.ProblemId, &submission.Status, &submission.Score, &submission.LanguageId, &submission.Length, &submission.Memory, &submission.Time, &submission.SourceCode, &createTimeStr, &updateTimeStr)
	log.Println(sql, id)
	if err != nil {
		return model.Submission{}, err
	}

	// 时间格式转换
	timeLayout := "2006-01-02 15:04:05"
	submission.CreateTime, err = time.Parse(timeLayout, createTimeStr)
	submission.UpdateTime, err = time.Parse(timeLayout, updateTimeStr)
	if err != nil {
		return model.Submission{}, err
	}

	return submission, nil
}

// 根据用户ID查询提交记录（不返回源代码）
func SelectSubmissionsByUserId(user_id uint64) ([]model.Submission, error) {
	sql := "SELECT id, problem_id, status, score, language_id, length, memory, time, create_time, update_time FROM tbl_submission WHERE user_id = ?"
	rows, err := Mysql.Query(sql, user_id)
	log.Println(sql, user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// 遍历查询结果
	submissions := make([]model.Submission, 0)
	for rows.Next() {
		var submission model.Submission
		var createTimeStr, updateTimeStr string

		submission.UserId = user_id

		err := rows.Scan(&submission.Id, &submission.ProblemId, &submission.Status, &submission.Score, &submission.LanguageId, &submission.Length, &submission.Memory, &submission.Time, &createTimeStr, &updateTimeStr)
		if err != nil {
			return nil, err
		}

		// 时间格式转换
		timeLayout := "2006-01-02 15:04:05"
		submission.CreateTime, err = time.Parse(timeLayout, createTimeStr)
		submission.UpdateTime, err = time.Parse(timeLayout, updateTimeStr)
		if err != nil {
			return nil, err
		}

		// 不返回源代码
		submission.SourceCode = ""

		//log.Println(submission)
		submissions = append(submissions, submission)
	}
	return submissions, nil
}

// 根据题目ID查询提交记录（不返回源代码）
func SelectSubmissionsByProblemId(problem_id uint64) ([]model.Submission, error) {
	sql := "SELECT id, user_id, status, score, language_id, length, memory, time, create_time, update_time FROM tbl_submission WHERE problem_id = ?"
	rows, err := Mysql.Query(sql, problem_id)
	log.Println(sql, problem_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// 遍历查询结果
	submissions := make([]model.Submission, 0)
	for rows.Next() {
		var submission model.Submission
		var createTimeStr, updateTimeStr string

		submission.ProblemId = problem_id

		err := rows.Scan(&submission.Id, &submission.UserId, &submission.Status, &submission.Score, &submission.LanguageId, &submission.Length, &submission.Memory, &submission.Time, &createTimeStr, &updateTimeStr)
		if err != nil {
			return nil, err
		}

		// 时间格式转换
		timeLayout := "2006-01-02 15:04:05"
		submission.CreateTime, err = time.Parse(timeLayout, createTimeStr)
		submission.UpdateTime, err = time.Parse(timeLayout, updateTimeStr)
		if err != nil {
			return nil, err
		}

		// 不返回源代码
		submission.SourceCode = ""

		//log.Println(submission)
		submissions = append(submissions, submission)
	}
	return submissions, nil
}

// 插入提交记录
func InsertSubmission(s model.Submission) (uint64, error) {
	sql := "INSERT INTO tbl_submission (user_id, problem_id, status, score, language_id, length, memory, time, source_code, create_time, update_time) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	stmt, err := Mysql.Prepare(sql)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	// 获取当前时间
	updateTime := time.Now().Format("2006-01-02 15:04:05")
	createTime := updateTime
	result, err := stmt.Exec(s.UserId, s.ProblemId, s.Status, s.Score, s.LanguageId, s.Length, s.Memory, s.Time, s.SourceCode, createTime, updateTime)
	log.Println(sql, s.UserId, s.ProblemId, s.Status, s.Score, s.LanguageId, s.Length, s.Memory, s.Time, s.SourceCode, createTime, updateTime)
	if err != nil {
		return 0, err
	}

	// 获取插入ID
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(id), nil
}

// 根据ID删除提交记录
func DeleteSubmissionById(id uint64) error {
	sql := "DELETE FROM tbl_submission WHERE id = ?"
	stmt, err := Mysql.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	log.Println(sql, id)
	if err != nil {
		return err
	}

	return nil
}

// 更新提交记录
func UpdateSubmissionById(s model.Submission) error {
	sql := "UPDATE tbl_submission SET user_id = ?, problem_id = ?, status = ?, score = ?, language_id = ?, length = ?, memory = ?, time = ?, source_code = ?, update_time = ? WHERE id = ?"
	stmt, err := Mysql.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// 获取当前时间
	updateTime := time.Now().Format("2006-01-02 15:04:05")
	_, err = stmt.Exec(s.UserId, s.ProblemId, s.Status, s.Score, s.LanguageId, s.Length, s.Memory, s.Time, s.SourceCode, updateTime, s.Id)
	log.Println(sql, s.UserId, s.ProblemId, s.Status, s.Score, s.LanguageId, s.Length, s.Memory, s.Time, s.SourceCode, updateTime, s.Id)
	if err != nil {
		return err
	}

	return nil
}
