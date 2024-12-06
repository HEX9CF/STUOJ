package comment

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"errors"
	"log"
)

func Select(condition dao.CommentWhere, userId uint64, page uint64, size uint64, admin ...bool) ([]entity.Comment, error) {
	comments, err := dao.SelectComments(condition, page, size)
	if err != nil {
		log.Println(err)
		return nil, errors.New("获取评论失败")
	}
	if len(admin) == 0 || !admin[0] {
		var publicComment []entity.Comment
		for _, comment := range comments {
			if comment.Status >= entity.CommentStatusPublic || comment.UserId == userId {
				publicComment = append(publicComment, comment)
			}
		}
		comments = publicComment
	}

	return comments, nil
}
