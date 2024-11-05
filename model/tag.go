package model

// 标签
type Tag struct {
	Id   uint64 `gorm:"primaryKey;autoIncrement;comment:标签ID" json:"id,omitempty"`
	Name string `gorm:"type:varchar(255);not null;unique;default:'';comment:标签名" json:"name,omitempty"`
}

func (Tag) TableName() string {
	return "tbl_tag"
}

// 题目标签关系
type ProblemTag struct {
	Id        uint64 `gorm:"primaryKey;autoIncrement;comment:关系ID" json:"id,omitempty"`
	ProblemId uint64 `gorm:"not null;default:0;comment:题目ID" json:"problem_id,omitempty"`
	TagId     uint64 `gorm:"not null;default:0;comment:标签ID" json:"tag_id,omitempty"`
}

func (ProblemTag) TableName() string {
	return "tbl_problem_tag"
}
