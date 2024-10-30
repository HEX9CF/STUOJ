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
	points := make([]model.Testcase, 0)
	for rows.Next() {
		var point model.Testcase

		point.ProblemId = problem_id

		err := rows.Scan(&point.Id, &point.Serial, &point.ProblemId, &point.TestInput, &point.TestOutput)
		if err != nil {
			return nil, err
		}

		//log.Println(point)
		points = append(points, point)
	}
	return points, nil
}
