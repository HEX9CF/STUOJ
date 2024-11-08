package dao

import (
	"STUOJ/internal/db"
	"STUOJ/internal/model"
)

// 添加评测点数据
func InsertTestcase(t model.Testcase) (uint64, error) {
	tx := db.Db.Create(&t)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return t.Id, nil
}

// 根据ID查询评测点数据
func SelectTestcaseById(id uint64) (model.Testcase, error) {
	var t model.Testcase

	tx := db.Db.Where("id = ?", id).First(&t)
	if tx.Error != nil {
		return model.Testcase{}, tx.Error
	}

	return t, nil
}

// 通过题目ID查询评测点数据
func SelectTestcasesByProblemId(problemId uint64) ([]model.Testcase, error) {
	var testcases []model.Testcase

	tx := db.Db.Where("problem_id = ?", problemId).Find(&testcases)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return testcases, nil
}

// 根据ID更新评测点数据
func UpdateTestcaseById(t model.Testcase) error {
	tx := db.Db.Model(&t).Where("id = ?", t.Id).Updates(map[string]interface{}{
		"serial":      t.Serial,
		"problem_id":  t.ProblemId,
		"test_input":  t.TestInput,
		"test_output": t.TestOutput,
	})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// 根据ID删除评测点数据
func DeleteTestcaseById(id uint64) error {
	tx := db.Db.Where("id = ?", id).Delete(&model.Testcase{})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
