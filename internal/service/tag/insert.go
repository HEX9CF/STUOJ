package tag

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"errors"
	"log"
)

// 插入标签
func Insert(t entity.Tag) (uint64, error) {
	var err error

	t.Id, err = dao.InsertTag(t)
	if err != nil {
		log.Println(err)
		return 0, errors.New("插入失败，标签名不能重复")
	}

	return t.Id, nil
}
