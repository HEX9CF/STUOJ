package blog

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"errors"
	"log"
)

// 根据ID查询博客
func SelectById(id uint64, userId uint64, admin ...bool) (entity.Blog, error) {
	b, err := dao.SelectBlogById(id)
	if err != nil {
		log.Println(err)
		return entity.Blog{}, errors.New("获取博客失败")
	}
	if b.Status != entity.BlogStatusPublic && (len(admin) == 0 || !admin[0]) {
		return entity.Blog{}, errors.New("该博客未公开")
	}
	return b, nil
}

func Select(condition dao.BlogWhere, userId uint64, page uint64, size uint64, admin ...bool) ([]entity.Blog, error) {
	blogs, err := dao.SelectBlogs(condition, page, size)
	if err != nil {
		log.Println(err)
		return nil, errors.New("获取博客失败")
	}
	if len(admin) == 0 || !admin[0] {
		var publicBlogs []entity.Blog
		for _, blog := range blogs {
			if blog.Status >= entity.BlogStatusPublic || blog.UserId == userId {
				publicBlogs = append(publicBlogs, blog)
			}
		}
		blogs = publicBlogs
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
