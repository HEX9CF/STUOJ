package dao

import (
	"STUOJ/internal/db"
	"STUOJ/internal/entity"
)

// 添加评测点数据
func InsertTestcase(t entity.Testcase) (uint64, error) {
	tx := db.Db.Create(&t)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return t.Id, nil
}

// 根据ID查询评测点数据
func SelectTestcaseById(id uint64) (entity.Testcase, error) {
	var t entity.Testcase

	tx := db.Db.Where("id = ?", id).First(&t)
	if tx.Error != nil {
		return entity.Testcase{}, tx.Error
	}

	return t, nil
}

// 通过题目ID查询评测点数据
func SelectTestcasesByProblemId(problemId uint64) ([]entity.Testcase, error) {
	var testcases []entity.Testcase

	tx := db.Db.Where("problem_id = ?", problemId).Find(&testcases)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return testcases, nil
}

// 根据ID更新评测点数据
func UpdateTestcaseById(t entity.Testcase) error {
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
	tx := db.Db.Where("id = ?", id).Delete(&entity.Testcase{})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// 根据题目ID删除评测点数据
func DeleteTestcasesByProblemId(pid uint64) error {
	tx := db.Db.Where("problem_id = ?", pid).Delete(&entity.Testcase{})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// 统计评测点数据数量
func CountTestcases() (uint64, error) {
	var count int64

	tx := db.Db.Model(&entity.Testcase{}).Count(&count)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return uint64(count), nil
}
