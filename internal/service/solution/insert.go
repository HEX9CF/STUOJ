package solution

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"errors"
	"log"
)

// 插入题解
func Insert(s entity.Solution) (uint64, error) {
	var err error

	// 插入题解
	s.Id, err = dao.InsertSolution(s)
	if err != nil {
		return 0, err
	}

	// 更新题目更新时间
	err = dao.UpdateProblemUpdateTimeById(s.ProblemId)
	if err != nil {
		log.Println(err)
		return 0, errors.New("更新题目更新时间失败")
	}

	return s.Id, nil
}
