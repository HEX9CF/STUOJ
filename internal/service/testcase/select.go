package testcase

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"errors"
	"log"
)

// 根据ID查询评测点数据
func SelectById(id uint64) (entity.Testcase, error) {
	t, err := dao.SelectTestcaseById(id)
	if err != nil {
		log.Println(err)
		return entity.Testcase{}, errors.New("查询失败")
	}

	return t, nil
}
