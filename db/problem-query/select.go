package problem_query

import (
	"STUOJ/db"
	"STUOJ/model"
	"log"
	"time"
)

// 根据ID查询题目
func SelectProblemById(id uint64) (model.Problem, error) {
	var p model.Problem
	var createTimeStr, updateTimeStr string

	p.Id = id

	sql := "SELECT title, source, difficulty, time_limit, memory_limit, description, input, output, sample_input, sample_output, hint, status, create_time, update_time FROM tbl_problem WHERE id = ? LIMIT 1"
	err := db.Mysql.QueryRow(sql, id).Scan(&p.Title, &p.Source, &p.Difficulty, &p.TimeLimit, &p.MemoryLimit, &p.Description, &p.Input, &p.Output, &p.SampleInput, &p.SampleOutput, &p.Hint, &p.Status, &createTimeStr, &updateTimeStr)
	log.Println(sql, id)
	if err != nil {
		return model.Problem{}, err
	}

	// 时间格式转换
	timeLayout := "2006-01-02 15:04:05"
	p.CreateTime, err = time.Parse(timeLayout, createTimeStr)
	if err != nil {
		return model.Problem{}, err
	}
	p.UpdateTime, err = time.Parse(timeLayout, updateTimeStr)
	if err != nil {
		return model.Problem{}, err
	}

	return p, nil
}

// 查询所有题目
func SelectAllProblems() ([]model.Problem, error) {
	sql := "SELECT id, title, source, difficulty, time_limit, memory_limit, description, input, output, sample_input, sample_output, hint, status, create_time, update_time FROM tbl_problem"
	rows, err := db.Mysql.Query(sql)
	log.Println(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// 遍历查询结果
	problems := make([]model.Problem, 0)
	for rows.Next() {
		var problem model.Problem
		var createTimeStr, updateTimeStr string

		err := rows.Scan(&problem.Id, &problem.Title, &problem.Source, &problem.Difficulty, &problem.TimeLimit, &problem.MemoryLimit, &problem.Description, &problem.Input, &problem.Output, &problem.SampleInput, &problem.SampleOutput, &problem.Hint, &problem.Status, &createTimeStr, &updateTimeStr)
		if err != nil {
			return nil, err
		}

		// 时间格式转换
		timeLayout := "2006-01-02 15:04:05"
		problem.CreateTime, err = time.Parse(timeLayout, createTimeStr)
		if err != nil {
			return nil, err
		}
		problem.UpdateTime, err = time.Parse(timeLayout, updateTimeStr)
		if err != nil {
			return nil, err
		}

		//log.Println(problem)
		problems = append(problems, problem)
	}
	return problems, nil
}

func SelectProblemsByTagId(tid uint64) ([]model.Problem, error) {
	sql := "SELECT id, title, source, difficulty, time_limit, memory_limit, description, input, output, sample_input, sample_output, hint, status, create_time, update_time FROM tbl_problem WHERE id IN (SELECT problem_id FROM tbl_problem_tag WHERE tag_id = ?)"
	rows, err := db.Mysql.Query(sql, tid)
	log.Println(sql, tid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// 遍历查询结果
	problems := make([]model.Problem, 0)
	for rows.Next() {
		var problem model.Problem
		var createTimeStr, updateTimeStr string

		err := rows.Scan(&problem.Id, &problem.Title, &problem.Source, &problem.Difficulty, &problem.TimeLimit, &problem.MemoryLimit, &problem.Description, &problem.Input, &problem.Output, &problem.SampleInput, &problem.SampleOutput, &problem.Hint, &problem.Status, &createTimeStr, &updateTimeStr)
		if err != nil {
			return nil, err
		}

		// 时间格式转换
		timeLayout := "2006-01-02 15:04:05"
		problem.CreateTime, err = time.Parse(timeLayout, createTimeStr)
		if err != nil {
			return nil, err
		}
		problem.UpdateTime, err = time.Parse(timeLayout, updateTimeStr)
		if err != nil {
			return nil, err
		}

		//log.Println(problem)
		problems = append(problems, problem)
	}
	return problems, nil
}

// 根据难度查询题目
func SelectProblemsByDifficulty(d model.ProblemDifficulty) ([]model.Problem, error) {
	sql := "SELECT id, title, source, difficulty, time_limit, memory_limit, description, input, output, sample_input, sample_output, hint, status, create_time, update_time FROM tbl_problem WHERE difficulty = ?"
	rows, err := db.Mysql.Query(sql, d)
	log.Println(sql, d)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// 遍历查询结果
	problems := make([]model.Problem, 0)
	for rows.Next() {
		var problem model.Problem
		var createTimeStr, updateTimeStr string

		err := rows.Scan(&problem.Id, &problem.Title, &problem.Source, &problem.Difficulty, &problem.TimeLimit, &problem.MemoryLimit, &problem.Description, &problem.Input, &problem.Output, &problem.SampleInput, &problem.SampleOutput, &problem.Hint, &problem.Status, &createTimeStr, &updateTimeStr)
		if err != nil {
			return nil, err
		}

		// 时间格式转换
		timeLayout := "2006-01-02 15:04:05"
		problem.CreateTime, err = time.Parse(timeLayout, createTimeStr)
		if err != nil {
			return nil, err
		}
		problem.UpdateTime, err = time.Parse(timeLayout, updateTimeStr)
		if err != nil {
			return nil, err
		}

		//log.Println(problem)
		problems = append(problems, problem)
	}
	return problems, nil
}
