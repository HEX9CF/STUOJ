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

// 根据ID删除博客（检查用户ID）
func DeleteByIdCheckUserId(id uint64, uid uint64, admin ...bool) error {
	// 查询博客
	b0, err := dao.SelectBlogById(id)
	if err != nil {
		log.Println(err)
		return errors.New("博客不存在")
	}

	// 检查权限
	if b0.UserId != uid && (len(admin) == 0 || !admin[0]) {
		return errors.New("没有权限，只能删除自己的博客")
	}

	// 删除博客
	err = dao.DeleteBlogById(id)
	if err != nil {
		log.Println(err)
		return errors.New("删除博客失败")
	}

	return nil
}
