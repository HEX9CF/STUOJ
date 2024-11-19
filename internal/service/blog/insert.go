package blog

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"errors"
	"log"
	"time"
)

// 插入博客
func Insert(b entity.Blog) (uint64, error) {
	var err error

	updateTime := time.Now()
	b.UpdateTime = updateTime
	b.CreateTime = updateTime

	// 插入博客
	b.Id, err = dao.InsertBlog(b)
	if err != nil {
		log.Println(err)
		return 0, errors.New("插入博客失败")
	}

	return b.Id, nil
}

// 保存草稿
func SaveDraft(b entity.Blog, uid uint64) (uint64, error) {
	var err error

	updateTime := time.Now()
	b.UpdateTime = updateTime
	b.CreateTime = updateTime

	b.Status = entity.BlogStatusDraft

	// 插入博客
	b.Id, err = dao.InsertBlog(b)
	if err != nil {
		log.Println(err)
		return 0, errors.New("插入博客失败")
	}

	return b.Id, nil
}
