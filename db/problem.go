package db

import (
	"STUOJ/model"
	"log"
	"time"
)

func GetProblemById(id uint64) (model.Problem, error) {
	// 查询题目
	var problem model.Problem
	var createTimeStr, updateTimeStr string
	sql := "SELECT id, title, source, difficulty, create_time, update_time FROM tbl_problem WHERE id = ? LIMIT 1"
	log.Println(sql)
	err := db.QueryRow(sql, id).Scan(&problem.Id, &problem.Title, &problem.Source, &problem.Difficulty, &createTimeStr, &updateTimeStr)
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

func GetAllProblems() ([]model.Problem, error) {
	// 查询所有题目
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

func SaveProblem(p model.Problem) error {
	// 插入题目
	sql := "INSERT INTO tbl_problem (title, source, difficulty, create_time, update_time) VALUES (?, ?, ?, ?, ?)"
	log.Println(sql)
	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	// 获取当前时间
	createTime := time.Now().Format("2006-01-02 15:04:05")
	updateTime := createTime
	_, err = stmt.Exec(p.Title, p.Source, p.Difficulty, createTime, updateTime)
	if err != nil {
		return err
	}

	return nil
}
