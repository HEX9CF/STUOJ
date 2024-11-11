package tag

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/model"
)

// 插入标签
func Insert(t model.Tag) (uint64, error) {
	var err error

	t.Id, err = dao.InsertTag(t)
	if err != nil {
		return 0, err
	}

	return t.Id, nil
}
