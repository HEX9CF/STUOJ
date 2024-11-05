package model

// 标签
type Tag struct {
	Id         uint64     `gorm:"primaryKey;autoIncrement;comment:标签ID" json:"id,omitempty"`
	Name       string     `gorm:"type:varchar(255);not null;unique;default:'';comment:标签名" json:"name,omitempty"`
	ProblemTag []*Problem `gorm:"many2many:tbl_problem_tag;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;association_jointable_foreignkey:problem_id;jointable_foreignkey:tag_id"`
}

func (Tag) TableName() string {
	return "tbl_tag"
}

// 题目标签关系
type ProblemTag struct {
	ProblemId uint64 `gorm:"primaryKey;autoIncrement:false;comment:题目ID"`
	TagId     uint64 `gorm:"primaryKey;autoIncrement:false;comment:标签ID"`
}

func (ProblemTag) TableName() string {
	return "tbl_problem_tag"
}
