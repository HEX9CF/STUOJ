package problem_query

import (
	"STUOJ/db"
	"STUOJ/model"
	"log"
	"time"
)

// 给题目添加标签
func InsertProblemTag(pid uint64, tid uint64) (uint64, error) {
	sql := "INSERT INTO tbl_problem_tag (problem_id, tag_id) VALUES (?, ?)"
	stmt, err := db.Mysql.Prepare(sql)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(pid, tid)
	log.Println(sql, pid, tid)
	if err != nil {
		return 0, err
	}

	// 获取插入ID
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// 更新题目更新时间
	err = UpdateProblemUpdateTimeById(pid)
	if err != nil {
		return uint64(id), err
	}

	return uint64(id), nil
}

// 查询题目标签关系是否存在
func CountProblemTagByProblemIdAndTagId(pid uint64, tid uint64) (uint64, error) {
	var count uint64
	sql := "SELECT COUNT(*) FROM tbl_problem_tag WHERE problem_id = ? AND tag_id = ?"
	err := db.Mysql.QueryRow(sql, pid, tid).Scan(&count)
	log.Println(sql, pid, tid)
	if err != nil {
		return 0, err
	}

	return count, nil
}

// 删除题目的某个标签
func DeleteProblemTagByProblemIdAndTagId(pid uint64, tid uint64) error {
	sql := "DELETE FROM tbl_problem_tag WHERE problem_id = ? AND tag_id = ?"
	stmt, err := db.Mysql.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(pid, tid)
	log.Println(sql, pid, tid)
	if err != nil {
		return err
	}

	// 更新题目更新时间
	err = UpdateProblemUpdateTimeById(pid)
	if err != nil {
		return err
	}

	return nil
}

func SelectTagsByProblemId(pid uint64) ([]model.Tag, error) {
	sql := "SELECT id, name FROM tbl_tag WHERE id IN (SELECT tag_id FROM tbl_problem_tag WHERE problem_id = ?)"
	rows, err := db.Mysql.Query(sql)
	log.Println(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// 遍历查询结果
	tags := make([]model.Tag, 0)
	for rows.Next() {
		var tag model.Tag

		err := rows.Scan(&tag.Id, &tag.Name)
		if err != nil {
			return nil, err
		}

		//log.Println(tag)
		tags = append(tags, tag)
	}
	return tags, nil
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
