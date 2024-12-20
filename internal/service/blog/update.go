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
		return errors.New("博客不存在")
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
		log.Println(err)
		return errors.New("更新博客失败")
	}

	return nil
}

// 用户发布博客
func SubmitByIdCheckUserId(id uint64, uid uint64, admin ...bool) error {
	// 查询博客
	b0, err := dao.SelectBlogById(id)
	if err != nil {
		log.Println(err)
		return errors.New("博客不存在")
	}

	// 检查权限
	if b0.UserId != uid {
		return errors.New("没有权限，只能提交自己的博客")
	}

	// // 检查博客状态
	// if b0.Status != entity.BlogStatusDraft {
	// 	return errors.New("博客不是草稿状态")
	// }

	updateTime := time.Now()
	b0.UpdateTime = updateTime
	if len(admin) > 0 && admin[0] {
		b0.Status = entity.BlogStatusPublic
	} else {
		b0.Status = entity.BLogStatusReview
	}
	// 更新博客
	err = dao.UpdateBlogById(b0)
	if err != nil {
		log.Println(err)
		return errors.New("更新博客失败")
	}

	return nil
}

// 用户编辑博客
func EditByIdCheckUserId(b entity.Blog, admin ...bool) error {
	// 查询博客
	b0, err := dao.SelectBlogById(b.Id)
	if err != nil {
		log.Println(err)
		return errors.New("博客不存在")
	}

	// 检查权限
	if b0.UserId != b.UserId && (len(admin) == 0 || !admin[0]) {
		return errors.New("没有权限，只能编辑自己的博客")
	}

	updateTime := time.Now()
	b0.Title = b.Title
	b0.Content = b.Content
	b0.UpdateTime = updateTime
	if len(admin) == 0 || !admin[0] || b0.Status == 0 {
		if (b0.Status > b.Status && b0.Status > entity.BLogStatusReview) || b0.Status == 0 {
			b0.Status = b.Status
		}
	}

	// 更新博客
	err = dao.UpdateBlogById(b0)
	if err != nil {
		log.Println(err)
		return errors.New("更新博客失败")
	}

	return nil
}
