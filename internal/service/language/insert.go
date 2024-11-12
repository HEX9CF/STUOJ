package language

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"errors"
	"log"
)

// 插入语言
func Insert(l entity.Language) (uint64, error) {
	var err error

	l.Id, err = dao.InsertLanguage(l)
	if err != nil {
		log.Println(err)
		return 0, errors.New("插入语言失败，语言名已存在")
	}

	return l.Id, nil
}
