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

	// 更新题目更新时间
	err = dao.UpdateProblemUpdateTimeById(t.ProblemId)
	if err != nil {
		log.Println(err)
		return 0, errors.New("更新题目更新时间失败")
	}

	// 插入评测点
	t.Id, err = dao.InsertTestcase(t)
	if err != nil {
		log.Println(err)
		return 0, errors.New("插入评测点失败")
	}

	return t.Id, nil
}
