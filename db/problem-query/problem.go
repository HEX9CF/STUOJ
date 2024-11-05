package problem_query

import (
	"STUOJ/db"
	"STUOJ/model"
	"log"
	"time"
)

// 插入题目
func InsertProblem(p model.Problem) (uint64, error) {
	sql := "INSERT INTO tbl_problem (title, source, difficulty, time_limit, memory_limit, description, input, output, sample_input, sample_output, hint, status, create_time, update_time) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	stmt, err := db.SqlDb.Prepare(sql)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	// 获取当前时间
	createTime := time.Now().Format("2006-01-02 15:04:05")
	updateTime := createTime
	result, err := stmt.Exec(p.Title, p.Source, p.Difficulty, p.TimeLimit, p.MemoryLimit, p.Description, p.Input, p.Output, p.SampleInput, p.SampleOutput, p.Hint, p.Status, createTime, updateTime)
	log.Println(sql, p.Title, p.Source, p.Difficulty, p.TimeLimit, p.MemoryLimit, p.Description, p.Input, p.Output, p.SampleInput, p.SampleOutput, p.Hint, p.Status, createTime, updateTime)
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

// 根据ID更新题目
func UpdateProblemById(p model.Problem) error {
	sql := "UPDATE tbl_problem SET title = ?, source = ?, difficulty = ?, time_limit = ?, memory_limit = ?, description = ?, input = ?, output = ?, sample_input = ?, sample_output = ?, hint = ?, status = ?, update_time = ? WHERE id = ?"
	stmt, err := db.SqlDb.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	// 获取当前时间
	updateTime := time.Now().Format("2006-01-02 15:04:05")
	_, err = stmt.Exec(p.Title, p.Source, p.Difficulty, p.TimeLimit, p.MemoryLimit, p.Description, p.Input, p.Output, p.SampleInput, p.SampleOutput, p.Hint, p.Status, updateTime, p.Id)
	log.Println(sql, p.Title, p.Source, p.Difficulty, p.TimeLimit, p.MemoryLimit, p.Description, p.Input, p.Output, p.SampleInput, p.SampleOutput, p.Hint, p.Status, updateTime, p.Id)
	if err != nil {
		return err
	}

	return nil
}

// 根据ID删除题目
func DeleteProblemById(id uint64) error {
	sql := "DELETE FROM tbl_problem WHERE id = ?"
	stmt, err := db.SqlDb.Prepare(sql)
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

// 根据ID更新提交记录状态更新时间
func UpdateProblemUpdateTimeById(id uint64) error {
	sql := "UPDATE tbl_problem SET update_time = ? WHERE id = ?"
	stmt, err := db.SqlDb.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	// 获取当前时间
	updateTime := time.Now().Format("2006-01-02 15:04:05")
	_, err = stmt.Exec(updateTime, id)
	log.Println(sql, updateTime, id)
	if err != nil {
		return err
	}

	return nil
}
