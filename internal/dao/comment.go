package dao

import (
	"STUOJ/internal/db"
	"STUOJ/internal/entity"
)

// 查询所有博客
func SelectAll() ([]entity.Comment, error) {
	var comments []entity.Comment

	tx := db.Db.Find(&comments)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return comments, nil
}

// 根据ID查询评论
func SelectCommentById(id uint64) (entity.Comment, error) {
	var c entity.Comment

	tx := db.Db.Where("id = ?", id).First(&c)
	if tx.Error != nil {
		return entity.Comment{}, tx.Error
	}

	return c, nil
}

// 查询所有评论
func SelectAllComments() ([]entity.Comment, error) {
	var comments []entity.Comment

	tx := db.Db.Find(&comments)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return comments, nil
}

// 根据博客ID查询评论
func SelectCommentsByBlogId(bid uint64) ([]entity.Comment, error) {
	var comments []entity.Comment

	tx := db.Db.Where("blog_id = ?", bid).Find(&comments)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return comments, nil
}

// 按状态查询评论
func SelectCommentsByStatus(s entity.CommentStatus) ([]entity.Comment, error) {
	var comments []entity.Comment

	tx := db.Db.Where("status = ?", s).Find(&comments)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return comments, nil
}

// 根据用户ID查询评论
func SelectCommentsByUserIdAndStatus(uid uint64, s entity.CommentStatus) ([]entity.Comment, error) {
	var comments []entity.Comment

	tx := db.Db.Where("user_id = ? AND status = ?", uid, s).Find(&comments)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return comments, nil
}

// 根据博客ID查询评论
func SelectCommentsByBlogIdAndStatus(bid uint64, s entity.CommentStatus) ([]entity.Comment, error) {
	var comments []entity.Comment

	tx := db.Db.Where("blog_id = ? AND status = ?", bid, s).Find(&comments)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return comments, nil
}

// 插入评论
func InsertComment(c entity.Comment) (uint64, error) {
	tx := db.Db.Create(&c)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return c.Id, nil
}

// 根据ID更新评论
func UpdateCommentById(b entity.Comment) error {
	tx := db.Db.Model(&b).Where("id = ?", b.Id).Updates(b)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// 根据ID删除评论
func DeleteCommentById(id uint64) error {
	tx := db.Db.Where("id = ?", id).Delete(&entity.Comment{})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// 统计评论数量
func CountComments() (int64, error) {
	var count int64

	tx := db.Db.Model(&entity.Comment{}).Count(&count)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return count, nil
}
