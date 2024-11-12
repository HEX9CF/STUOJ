package testcase

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"errors"
	"log"
)

// 添加评测点数据
func Insert(t entity.Testcase) (uint64, error) {
	var err error

	t.Id, err = dao.InsertTestcase(t)
	if err != nil {
		log.Println(err)
		return 0, errors.New("插入评测点失败")
	}

	return t.Id, nil
}
