package problem_query

import (
	"STUOJ/db"
	"STUOJ/model"
	"log"
)

// 根据ID查询评测点数据
func SelectTestcaseById(id uint64) (model.Testcase, error) {
	var t model.Testcase

	t.Id = id

	sql := "SELECT serial, problem_id, test_input, test_output FROM tbl_testcase WHERE id = ?"
	err := db.Mysql.QueryRow(sql, id).Scan(&t.Serial, &t.ProblemId, &t.TestInput, &t.TestOutput)
	log.Println(sql, id)
	if err != nil {
		return model.Testcase{}, err
	}

	return t, nil
}

// 通过题目ID查询评测点数据
func SelectTestcasesByProblemId(problem_id uint64) ([]model.Testcase, error) {
	sql := "SELECT id, serial, problem_id, test_input, test_output FROM tbl_testcase WHERE problem_id = ?"
	rows, err := db.Mysql.Query(sql, problem_id)
	log.Println(sql, problem_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// 遍历查询结果
	testcases := make([]model.Testcase, 0)
	for rows.Next() {
		var testcase model.Testcase

		testcase.ProblemId = problem_id

		err := rows.Scan(&testcase.Id, &testcase.Serial, &testcase.ProblemId, &testcase.TestInput, &testcase.TestOutput)
		if err != nil {
			return nil, err
		}

		//log.Println(testcase)
		testcases = append(testcases, testcase)
	}
	return testcases, nil
}

// 添加评测点数据
func InsertTestcase(t model.Testcase) (uint64, error) {
	sql := "INSERT INTO tbl_testcase (serial, problem_id, test_input, test_output) VALUES (?, ?, ?, ?)"
	stmt, err := db.Mysql.Prepare(sql)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(t.Serial, t.ProblemId, t.TestInput, t.TestOutput)
	log.Println(sql, t.Serial, t.ProblemId, t.TestInput, t.TestOutput)
	if err != nil {
		return 0, err
	}

	// 获取插入ID
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// 更新题目更新时间
	err = UpdateProblemUpdateTimeById(t.ProblemId)
	if err != nil {
		return uint64(id), err
	}

	return uint64(id), nil
}

// 根据ID更新评测点数据
func UpdateTestcaseById(t model.Testcase) error {
	sql := "UPDATE tbl_testcase SET serial = ?, problem_id = ?, test_input = ?, test_output = ? WHERE id = ?"
	stmt, err := db.Mysql.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(t.Serial, t.ProblemId, t.TestInput, t.TestOutput, t.Id)
	log.Println(sql, t.Serial, t.ProblemId, t.TestInput, t.TestOutput, t.Id)
	if err != nil {
		return err
	}

	// 更新题目更新时间
	err = UpdateProblemUpdateTimeById(t.ProblemId)
	if err != nil {
		return err
	}

	return nil
}

// 根据ID删除评测点数据
func DeleteTestcaseById(id uint64) error {
	sql := "DELETE FROM tbl_testcase WHERE id = ?"
	stmt, err := db.Mysql.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	log.Println(sql, id)
	if err != nil {
		return err
	}

	// 更新题目更新时间
	err = UpdateProblemUpdateTimeById(id)
	if err != nil {
		return err
	}

	return nil
}
