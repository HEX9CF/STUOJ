package model

import "STUOJ/internal/entity"

// 博客数据（博客+标签+评论）
type BlogData struct {
	Blog entity.Blog  `json:"blog"`
	Tags []entity.Tag `json:"tags,omitempty"`
}
