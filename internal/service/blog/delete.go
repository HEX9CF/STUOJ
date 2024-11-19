package blog

import (
	"STUOJ/internal/dao"
	"errors"
	"log"
)

// 根据ID删除博客
func DeleteById(id uint64) error {
	// 查询博客
	_, err := dao.SelectBlogById(id)
	if err != nil {
		log.Println(err)
		return errors.New("博客不存在")
	}

	// 删除博客
	err = dao.DeleteBlogById(id)
	if err != nil {
		log.Println(err)
		return errors.New("删除博客失败")
	}

	return nil
}
