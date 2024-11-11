package solution

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/model"
)

// 插入题解
func Insert(s model.Solution) (uint64, error) {
	var err error

	s.Id, err = dao.InsertSolution(s)
	if err != nil {
		return 0, err
	}

	return s.Id, nil
}
