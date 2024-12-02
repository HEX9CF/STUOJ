package comment

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"errors"
	"log"
)

func SelectAll() ([]entity.Comment, error) {
	comments, err := dao.SelectAllComments()
	if err != nil {
		log.Println(err)
		return nil, errors.New("获取评论失败")
	}

	return comments, nil
}

// 根据用户ID查询公开评论
func SelectPublicByUserId(uid uint64) ([]entity.Comment, error) {
	comments, err := dao.SelectCommentsByUserIdAndStatus(uid, entity.CommentStatusPublic)
	if err != nil {
		log.Println(err)
		return nil, errors.New("获取评论失败")
	}

	return comments, nil
}

// 根据博客ID查询公开评论
func SelectPublicByBlogId(bid uint64) ([]entity.Comment, error) {
	comments, err := dao.SelectCommentsByBlogIdAndStatus(bid, entity.CommentStatusPublic)
	if err != nil {
		log.Println(err)
		return nil, errors.New("获取评论失败")
	}

	return comments, nil
}
