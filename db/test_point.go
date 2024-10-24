package db

import (
	"STUOJ/model"
	"log"
)

// 根据题目ID查询提交记录
func SelectTestPointsByProblemId(problem_id uint64) ([]model.TestPoint, error) {
	sql := "SELECT id, serial, problem_id, test_input, test_output FROM tbl_test_point WHERE problem_id = ?"
	rows, err := db.Query(sql, problem_id)
	log.Println(sql, problem_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// 遍历查询结果
	points := make([]model.TestPoint, 0)
	for rows.Next() {
		var point model.TestPoint

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
