package problem_query

import (
	"STUOJ/db"
	"log"
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
