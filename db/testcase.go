package db

import (
	"STUOJ/model"
	"log"
)

// 通过题目ID查询评测点数据
func SelectTestcasesByProblemId(problem_id uint64) ([]model.Testcase, error) {
	sql := "SELECT id, serial, problem_id, test_input, test_output FROM tbl_testcase WHERE problem_id = ?"
	rows, err := Mysql.Query(sql, problem_id)
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
