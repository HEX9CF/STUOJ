package blog

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"errors"
	"log"
	"time"
)

// 根据ID更新博客
func UpdateById(b entity.Blog) error {
	// 查询博客
	b0, err := dao.SelectBlogById(b.Id)
	if err != nil {
		log.Println(err)
		return errors.New("获取博客失败")
	}

	updateTime := time.Now()
	b0.UpdateTime = updateTime
	b0.UserId = b.UserId
	b0.ProblemId = b.ProblemId
	b0.Title = b.Title
	b0.Content = b.Content
	b0.Status = b.Status

	// 更新博客
	err = dao.UpdateBlogById(b0)
	if err != nil {
		return err
	}

	return nil
}

// 用户发布博客（变为待审核状态）
func Publish(id uint64, uid uint64) error {
	// 查询博客
	b0, err := dao.SelectBlogById(id)
	if err != nil {
		log.Println(err)
		return errors.New("获取博客失败")
	}

	if b0.UserId != uid {
		return errors.New("没有权限")
	}

	updateTime := time.Now()
	b0.UpdateTime = updateTime
	b0.Status = entity.BLogStatusReview

	// 更新博客
	err = dao.UpdateBlogById(b0)
	if err != nil {
		log.Println(err)
		return errors.New("更新博客失败")
	}

	return nil
}

// 用户编辑博客（变回草稿状态）
func Edit(b entity.Blog, uid uint64) error {
	// 查询博客
	b0, err := dao.SelectBlogById(b.Id)
	if err != nil {
		log.Println(err)
		return errors.New("获取博客失败")
	}

	if b0.UserId != uid {
		return errors.New("没有权限")
	}

	updateTime := time.Now()
	b0.Title = b.Title
	b0.Content = b.Content
	b0.UpdateTime = updateTime
	b0.Status = entity.BLogStatusReview

	// 更新博客
	err = dao.UpdateBlogById(b0)
	if err != nil {
		log.Println(err)
		return errors.New("更新博客失败")
	}

	return nil
}

// 博客通过审核
func Approve(id uint64) error {
	// 查询博客
	b0, err := dao.SelectBlogById(id)
	if err != nil {
		log.Println(err)
		return errors.New("获取博客失败")
	}

	updateTime := time.Now()
	b0.UpdateTime = updateTime
	b0.Status = entity.BlogStatusPublic

	// 更新博客
	err = dao.UpdateBlogById(b0)
	if err != nil {
		log.Println(err)
		return errors.New("更新博客失败")
	}

	return nil
}

// 封禁博客
func Ban(id uint64) error {
	// 查询博客
	b0, err := dao.SelectBlogById(id)
	if err != nil {
		log.Println(err)
		return errors.New("获取博客失败")
	}

	updateTime := time.Now()
	b0.UpdateTime = updateTime
	b0.Status = entity.BLogStatusBanned

	// 更新博客
	err = dao.UpdateBlogById(b0)
	if err != nil {
		log.Println(err)
		return errors.New("更新博客失败")
	}

	return nil
}
