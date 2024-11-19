package blog

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"errors"
	"log"
)

// 根据ID查询博客
func SelectById(id uint64) (entity.Blog, error) {
	b, err := dao.SelectBlogById(id)
	if err != nil {
		return entity.Blog{}, errors.New("获取博客失败")
	}

	return b, nil
}

// 根据ID查询公开博客
func SelectPublicById(id uint64) (entity.Blog, error) {
	b, err := dao.SelectBlogByIdAndStatus(id, entity.BlogStatusPublic)
	if err != nil {
		return entity.Blog{}, errors.New("获取博客失败")
	}

	return b, nil
}

// 查询所有博客
func SelectAll() ([]entity.Blog, error) {
	blogs, err := dao.SelectAllBlogs()
	if err != nil {
		return nil, err
	}

	hideBlogContent(blogs)

	return blogs, nil
}

// 查询公开博客
func SelectPublic() ([]entity.Blog, error) {
	blogs, err := dao.SelectBlogsByStatus(entity.BlogStatusPublic)
	if err != nil {
		return nil, err
	}

	hideBlogContent(blogs)

	return blogs, nil
}

// 根据用户ID查询博客
func SelectPublicByUserId(uid uint64) ([]entity.Blog, error) {
	blogs, err := dao.SelectBlogsByUserIdAndStatus(uid, entity.BlogStatusPublic)
	if err != nil {
		return nil, err
	}

	hideBlogContent(blogs)

	return blogs, nil
}

// 根据题目ID查询博客
func SelectPublicByProblemId(pid uint64) ([]entity.Blog, error) {
	blogs, err := dao.SelectBlogsByProblemIdAndStatus(pid, entity.BlogStatusPublic)
	if err != nil {
		return nil, err
	}

	hideBlogContent(blogs)

	return blogs, nil
}

// 根据状态查询并根据标题模糊查询公开博客
func SelectPublicLikeTitle(title string) ([]entity.Blog, error) {
	var blogs []entity.Blog

	blogs, err := dao.SelectBlogsLikeTitleByStatus(title, entity.BlogStatusPublic)
	if err != nil {
		log.Println(err)
		return nil, errors.New("获取博客失败")
	}

	hideBlogContent(blogs)

	return blogs, nil
}

// 不返回正文
func hideBlogContent(blogs []entity.Blog) {
	for i := range blogs {
		blogs[i].Content = ""
	}
}
