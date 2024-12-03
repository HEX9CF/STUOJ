package dao

import (
	"STUOJ/internal/db"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"time"
)

// 插入博客
func InsertBlog(b entity.Blog) (uint64, error) {
	tx := db.Db.Create(&b)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return b.Id, nil
}

// 根据ID查询博客
func SelectBlogById(id uint64) (entity.Blog, error) {
	var b entity.Blog

	tx := db.Db.Where("id = ?", id).First(&b)
	if tx.Error != nil {
		return entity.Blog{}, tx.Error
	}

	return b, nil
}

// 根据ID查询博客
func SelectBlogByIdAndStatus(id uint64, s entity.BlogStatus) (entity.Blog, error) {
	var b entity.Blog

	tx := db.Db.Where("id = ? AND status = ?", id, s).First(&b)
	if tx.Error != nil {
		return entity.Blog{}, tx.Error
	}

	return b, nil
}

// 查询所有博客
func SelectAllBlogs() ([]entity.Blog, error) {
	var blogs []entity.Blog

	tx := db.Db.Find(&blogs)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return blogs, nil
}

// 按状态查询博客
func SelectBlogsByStatus(s entity.BlogStatus) ([]entity.Blog, error) {
	var blogs []entity.Blog

	tx := db.Db.Where("status = ?", s).Find(&blogs)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return blogs, nil
}

// 根据用户ID查询博客
func SelectBlogsByUserIdAndStatus(uid uint64, s entity.BlogStatus) ([]entity.Blog, error) {
	var blogs []entity.Blog

	tx := db.Db.Where("user_id = ? AND status = ?", uid, s).Find(&blogs)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return blogs, nil
}

// 根据题目ID查询博客
func SelectBlogsByProblemIdAndStatus(pid uint64, s entity.BlogStatus) ([]entity.Blog, error) {
	var blogs []entity.Blog

	tx := db.Db.Where("problem_id = ? AND status = ?", pid, s).Find(&blogs)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return blogs, nil
}

// 根据状态查询并根据标题模糊查询博客
func SelectBlogsLikeTitleByStatus(title string, s entity.BlogStatus) ([]entity.Blog, error) {
	var blogs []entity.Blog

	tx := db.Db.Where("status = ? AND title like ?", s, "%"+title+"%").Find(&blogs)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return blogs, nil
}

// 根据ID更新博客
func UpdateBlogById(b entity.Blog) error {
	tx := db.Db.Model(&b).Where("id = ?", b.Id).Updates(b)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// 根据ID删除博客
func DeleteBlogById(id uint64) error {
	tx := db.Db.Where("id = ?", id).Delete(&entity.Blog{})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// 统计博客数量
func CountBlogs() (uint64, error) {
	var count int64

	tx := db.Db.Model(&entity.Blog{}).Count(&count)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return uint64(count), nil
}

// 根据创建时间统计博客数量
func CountBlogsBetweenCreateTime(startTime time.Time, endTime time.Time) ([]model.CountByDate, error) {
	var counts []model.CountByDate

	tx := db.Db.Model(&entity.Blog{}).Where("create_time between ? and ?", startTime, endTime).Select("date(create_time) as date, count(*) as count").Group("date").Scan(&counts)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return counts, nil
}
