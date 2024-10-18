package db

import (
	"STUOJ/model"
	"log"
	"time"
)

// 根据ID查询题目
func SelectProblemById(id uint64) (model.Problem, error) {
	var problem model.Problem
	var createTimeStr, updateTimeStr string
	sql := "SELECT title, source, difficulty, create_time, update_time FROM tbl_problem WHERE id = ? LIMIT 1"
	err := db.QueryRow(sql, id).Scan(&problem.Title, &problem.Source, &problem.Difficulty, &createTimeStr, &updateTimeStr)
	log.Println(sql, id)
	if err != nil {
		return model.Problem{}, err
	}

	// 时间格式转换
	timeLayout := "2006-01-02 15:04:05"
	problem.CreateTime, err = time.Parse(timeLayout, createTimeStr)
	if err != nil {
		return model.Problem{}, err
	}
	problem.UpdateTime, err = time.Parse(timeLayout, updateTimeStr)
	if err != nil {
		return model.Problem{}, err
	}

	return problem, nil
}

// 查询所有题目
func SelectAllProblems() ([]model.Problem, error) {
	sql := "SELECT id, title, source, difficulty, create_time, update_time FROM tbl_problem"
	rows, err := db.Query(sql)
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

		err := rows.Scan(&problem.Id, &problem.Title, &problem.Source, &problem.Difficulty, &createTimeStr, &updateTimeStr)
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

// 插入题目
func InsertProblem(p model.Problem) error {
	sql := "INSERT INTO tbl_problem(title, source, difficulty, create_time, update_time) VALUES(?, ?, ?, ?, ?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	// 获取当前时间
	createTime := time.Now().Format("2006-01-02 15:04:05")
	updateTime := createTime
	_, err = stmt.Exec(p.Title, p.Source, p.Difficulty, createTime, updateTime)
	log.Println(sql, p.Title, p.Source, p.Difficulty, createTime, updateTime)
	if err != nil {
		return err
	}

	return nil
}

// 根据ID更新题目
func UpdateProblemById(p model.Problem) error {
	sql := "UPDATE tbl_problem SET title = ?, source = ?, difficulty = ?, update_time = ? WHERE id = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	// 获取当前时间
	updateTime := time.Now().Format("2006-01-02 15:04:05")
	_, err = stmt.Exec(p.Title, p.Source, p.Difficulty, updateTime, p.Id)
	log.Println(sql, p.Title, p.Source, p.Difficulty, updateTime, p.Id)
	if err != nil {
		return err
	}

	return nil
}

// 根据ID删除题目
func DeleteProblemById(id uint64) error {
	sql := "DELETE FROM tbl_problem WHERE id = ?"
	stmt, err := db.Prepare(sql)
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
