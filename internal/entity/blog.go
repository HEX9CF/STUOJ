package entity

import "time"

// 博客
type Blog struct {
	Id         uint      `gorm:"primaryKey;autoIncrement;comment:'博客ID'"`
	UserId     uint      `gorm:"not null;default:0;comment:'用户ID'"`
	ProblemId  uint      `gorm:"not null;default:0;comment:'关联题目ID：0 不关联'"`
	Title      string    `gorm:"type:text;not null;comment:'标题'"`
	Content    string    `gorm:"type:longtext;not null;comment:'内容'"`
	Status     uint      `gorm:"not null;default:1;comment:'状态：0 被屏蔽，1 公开，2 草稿'"`
	CreateTime time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:'创建时间'"`
	UpdateTime time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP;comment:'更新时间'"`
}

func (Blog) TableName() string {
	return "tbl_blog"
}
