package model

import "STUOJ/internal/entity"

// 博客数据（博客+评论）
type BlogData struct {
	Blog     entity.Blog      `json:"blog,omitempty"`
	Comments []entity.Comment `json:"comments,omitempty"`
}
