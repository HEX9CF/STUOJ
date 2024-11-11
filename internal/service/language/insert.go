package language

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/model"
)

// 插入语言
func Insert(l model.Language) (uint64, error) {
	var err error

	l.Id, err = dao.InsertLanguage(l)
	if err != nil {
		return 0, err
	}

	return l.Id, nil
}
