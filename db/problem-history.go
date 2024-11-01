package db

import (
	"STUOJ/model"
	"log"
	"time"
)

// 根据题目ID查询题目历史记录
func SelectProblemHistoriesByProblemId(pid uint64) ([]model.ProblemHistory, error) {
	sql := "SELECT id, user_id, problem_id, title, source, difficulty, time_limit, memory_limit, description, input, output, sample_input, sample_output, hint, operation, create_time FROM tbl_problem_history WHERE problem_id = ?"
	rows, err := Mysql.Query(sql)
	log.Println(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// 遍历查询结果
	phs := make([]model.ProblemHistory, 0)
	for rows.Next() {
		var ph model.ProblemHistory
		var createTimeStr string

		err := rows.Scan(&ph.Id, &ph.UserId, &ph.ProblemId, &ph.Title, &ph.Source, &ph.Difficulty, &ph.TimeLimit, &ph.MemoryLimit, &ph.Description, &ph.Input, &ph.Output, &ph.SampleInput, &ph.SampleOutput, &ph.Hint, &ph.Operation, &createTimeStr)
		if err != nil {
			return nil, err
		}

		// 时间格式转换
		timeLayout := "2006-01-02 15:04:05"
		ph.CreateTime, err = time.Parse(timeLayout, createTimeStr)
		if err != nil {
			return nil, err
		}

		//log.Println(ph)
		phs = append(phs, ph)
	}
	return phs, nil
}

// 插入题目历史记录
func InsertProblemHistory(p model.Problem, uid uint64, op model.Operation) (uint64, error) {
	sql := "INSERT INTO tbl_problem_history (user_id, problem_id, title, source, difficulty, time_limit, memory_limit, description, input, output, sample_input, sample_output, hint, operation, create_time) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	stmt, err := Mysql.Prepare(sql)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	// 获取当前时间
	createTime := time.Now().Format("2006-01-02 15:04:05")
	result, err := stmt.Exec(uid, p.Id, p.Title, p.Source, p.Difficulty, p.TimeLimit, p.MemoryLimit, p.Description, p.Input, p.Output, p.SampleInput, p.SampleOutput, p.Hint, op, createTime)
	log.Println(sql, uid, p.Id, p.Title, p.Source, p.Difficulty, p.TimeLimit, p.MemoryLimit, p.Description, p.Input, p.Output, p.SampleInput, p.SampleOutput, p.Hint, op, createTime)
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

// 根据ID删除题目历史记录
func DeleteProblemHistoryById(id uint64) error {
	sql := "DELETE FROM tbl_problem_history WHERE id = ?"
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
