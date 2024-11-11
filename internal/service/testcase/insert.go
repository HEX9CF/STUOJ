package testcase

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/model"
)

// 添加评测点数据
func Insert(t model.Testcase) (uint64, error) {
	var err error

	t.Id, err = dao.InsertTestcase(t)
	if err != nil {
		return 0, err
	}

	return t.Id, nil
}
