package blog

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"errors"
	"log"
)

// 根据ID查询博客
func SelectById(id uint64) (model.BlogData, error) {
	b, err := dao.SelectBlogById(id)
	if err != nil {
		log.Println(err)
		return model.BlogData{}, errors.New("获取博客失败")
	}

	// 查询博客评论
	comments, err := dao.SelectCommentsByBlogId(id)
	if err != nil {
		log.Println(err)
		return model.BlogData{}, errors.New("获取博客评论失败")
	}

	// 封装博客数据
	bd := model.BlogData{
		Blog:     b,
		Comments: comments,
	}

	return bd, nil
}

// 根据ID查询公开博客
func SelectPublicById(id uint64) (model.BlogData, error) {
	b, err := dao.SelectBlogByIdAndStatus(id, entity.BlogStatusPublic)
	if err != nil {
		log.Println(err)
		return model.BlogData{}, errors.New("获取博客失败")
	}

	// 查询博客评论
	comments, err := dao.SelectCommentsByBlogIdAndStatus(id, entity.CommentStatusPublic)
	if err != nil {
		log.Println(err)
		return model.BlogData{}, errors.New("获取博客评论失败")
	}

	bd := model.BlogData{
		Blog:     b,
		Comments: comments,
	}

	return bd, nil
}

// 查询所有博客
func SelectAll() ([]model.BlogData, error) {
	blogs, err := dao.SelectAllBlogs()
	if err != nil {
		log.Println(err)
		return nil, errors.New("获取博客失败")
	}

	bds := wrapBlogDatas(blogs)

	return bds, nil
}

// 查询公开博客
func SelectPublic() ([]model.BlogData, error) {
	blogs, err := dao.SelectBlogsByStatus(entity.BlogStatusPublic)
	if err != nil {
		log.Println(err)
		return nil, errors.New("获取博客失败")
	}

	bds := wrapBlogDatas(blogs)

	return bds, nil
}

// 根据用户ID查询公开博客
func SelectPublicByUserId(uid uint64) ([]model.BlogData, error) {
	blogs, err := dao.SelectBlogsByUserIdAndStatus(uid, entity.BlogStatusPublic)
	if err != nil {
		log.Println(err)
		return nil, errors.New("获取博客失败")
	}

	bds := wrapBlogDatas(blogs)

	return bds, nil
}

// 根据用户ID查询草稿箱
func SelectDraftByUserId(uid uint64) ([]model.BlogData, error) {
	blogs, err := dao.SelectBlogsByUserIdAndStatus(uid, entity.BlogStatusDraft)
	if err != nil {
		log.Println(err)
		return nil, errors.New("获取博客失败")
	}

	bds := wrapBlogDatas(blogs)

	return bds, nil
}

// 根据题目ID查询博客
func SelectPublicByProblemId(pid uint64) ([]model.BlogData, error) {
	blogs, err := dao.SelectBlogsByProblemIdAndStatus(pid, entity.BlogStatusPublic)
	if err != nil {
		log.Println(err)
		return nil, errors.New("获取博客失败")
	}

	bds := wrapBlogDatas(blogs)

	return bds, nil
}

// 根据状态查询并根据标题模糊查询公开博客
func SelectPublicLikeTitle(title string) ([]model.BlogData, error) {
	blogs, err := dao.SelectBlogsLikeTitleByStatus(title, entity.BlogStatusPublic)
	if err != nil {
		log.Println(err)
		return nil, errors.New("获取博客失败")
	}

	bds := wrapBlogDatas(blogs)

	return bds, nil
}

// 不返回正文
func hideBlogContent(blogs []entity.Blog) {
	for i := range blogs {
		blogs[i].Content = ""
	}
}

// 封装博客数据
func wrapBlogDatas(blogs []entity.Blog) []model.BlogData {
	var bds []model.BlogData

	hideBlogContent(blogs)

	for _, b := range blogs {
		bd := model.BlogData{
			Blog: b,
		}

		bds = append(bds, bd)
	}

	return bds
}
