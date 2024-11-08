package model

import (
	"STUOJ/utils"
	"time"

	"golang.org/x/exp/rand"
)

// 状态：0 作废，1 公开，2 出题中，3 调试中
type ProblemStatus uint8

const (
	ProblemStatusInvalid   ProblemStatus = 0
	ProblemStatusPublic    ProblemStatus = 1
	ProblemStatusEditing   ProblemStatus = 2
	ProblemStatusDebugging ProblemStatus = 3
)

func (s ProblemStatus) String() string {
	switch s {
	case ProblemStatusInvalid:
		return "作废"
	case ProblemStatusPublic:
		return "公开"
	case ProblemStatusEditing:
		return "出题中"
	case ProblemStatusDebugging:
		return "调试中"
	default:
		return "未知状态"
	}
}

// 难度： 0 暂无评定，1 普及−，2 普及/提高−，3 普及+/提高，4 提高+/省选− ，5 省选/NOI−，6 NOI/NOI+/CTSC
type ProblemDifficulty uint8

const (
	ProblemDifficultyUnknown ProblemDifficulty = 0
	ProblemDifficulty1       ProblemDifficulty = 1
	ProblemDifficulty2       ProblemDifficulty = 2
	ProblemDifficulty3       ProblemDifficulty = 3
	ProblemDifficulty4       ProblemDifficulty = 4
	ProblemDifficulty5       ProblemDifficulty = 5
	ProblemDifficulty6       ProblemDifficulty = 6
)

func (d ProblemDifficulty) String() string {
	switch d {
	case ProblemDifficultyUnknown:
		return "暂无评定"
	case ProblemDifficulty1:
		return "普及−"
	case ProblemDifficulty2:
		return "普及/提高−"
	case ProblemDifficulty3:
		return "普及+/提高"
	case ProblemDifficulty4:
		return "提高+/省选−"
	case ProblemDifficulty5:
		return "省选/NOI−"
	case ProblemDifficulty6:
		return "NOI/NOI+/CTSC"
	default:
		return "暂无评定"
	}
}

// 题目
type Problem struct {
	Id           uint64            `gorm:"primaryKey;autoIncrement;comment:题目ID" json:"id,omitempty"`
	Title        string            `gorm:"type:text;not null;comment:标题" json:"title,omitempty"`
	Source       string            `gorm:"type:text;not null;comment:题目来源" json:"source,omitempty"`
	Difficulty   ProblemDifficulty `gorm:"not null;default:0;comment:难度" json:"difficulty,omitempty"`
	TimeLimit    float64           `gorm:"not null;default:1;comment:时间限制（s）" json:"time_limit,omitempty"`
	MemoryLimit  uint64            `gorm:"not null;default:131072;comment:内存限制（kb）" json:"memory_limit,omitempty"`
	Description  string            `gorm:"type:longtext;not null;comment:题面" json:"description,omitempty"`
	Input        string            `gorm:"type:longtext;not null;comment:输入说明" json:"input,omitempty"`
	Output       string            `gorm:"type:longtext;not null;comment:输出说明" json:"output,omitempty"`
	SampleInput  string            `gorm:"type:longtext;not null;comment:输入样例" json:"sample_input,omitempty"`
	SampleOutput string            `gorm:"type:longtext;not null;comment:输出样例" json:"sample_output,omitempty"`
	Hint         string            `gorm:"type:longtext;not null;comment:提示" json:"hint,omitempty"`
	Status       ProblemStatus     `gorm:"not null;default:1;comment:状态" json:"status,omitempty"`
	ProblemTag   []*Tag            `gorm:"many2many:tbl_problem_tag;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;association_jointable_foreignkey:tag_id;jointable_foreignkey:problem_id" json:"problem_tag,omitempty"`
	CreateTime   time.Time         `gorm:"not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time,omitempty"`
	UpdateTime   time.Time         `gorm:"not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"update_time,omitempty"`
}

func (Problem) TableName() string {
	return "tbl_problem"
}

// 题解
type Solution struct {
	Id         uint64 `gorm:"primaryKey;autoIncrement;comment:题解ID" json:"id,omitempty"`
	ProblemId  uint64 `gorm:"not null;default:0;comment:题目ID" json:"problem_id,omitempty"`
	LanguageId uint64 `gorm:"not null;default:0;comment:语言ID" json:"language_id,omitempty"`
	SourceCode string `gorm:"type:longtext;not null;comment:源代码" json:"source_code,omitempty"`
}

func (Solution) TableName() string {
	return "tbl_solution"
}

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

// 评测点数据
type Testcase struct {
	Id         uint64 `gorm:"primaryKey;autoIncrement;comment:评测点ID" json:"id,omitempty"`
	Serial     uint64 `gorm:"not null;default:0;comment:评测点序号" json:"serial,omitempty"`
	ProblemId  uint64 `gorm:"not null;default:0;comment:题目ID" json:"problem_id,omitempty"`
	TestInput  string `gorm:"type:longtext;not null;comment:测试输入" json:"test_input,omitempty"`
	TestOutput string `gorm:"type:longtext;not null;comment:测试输出" json:"test_output,omitempty"`
}

func (Testcase) TableName() string {
	return "tbl_testcase"
}

// 操作：0 未知，1 添加，2 修改，3 删除
type Operation uint8

const (
	OperationUnknown Operation = 0
	OperationAdd     Operation = 1
	OperationUpdate  Operation = 2
	OperationDelete  Operation = 3
)

func (o Operation) String() string {
	switch o {
	case OperationUnknown:
		return "未知"
	case OperationAdd:
		return "添加"
	case OperationUpdate:
		return "修改"
	case OperationDelete:
		return "删除"
	default:
		return "未知"
	}
}

// 题目历史记录
type ProblemHistory struct {
	Id           uint64            `gorm:"primaryKey;autoIncrement;comment:记录ID" json:"id,omitempty"`
	UserId       uint64            `gorm:"not null;default:0;comment:用户ID" json:"user_id,omitempty"`
	ProblemId    uint64            `gorm:"not null;default:0;comment:题目ID" json:"problem_id,omitempty"`
	Title        string            `gorm:"type:text;not null;comment:标题" json:"title,omitempty"`
	Source       string            `gorm:"type:text;not null;comment:题目来源" json:"source,omitempty"`
	Difficulty   ProblemDifficulty `gorm:"not null;default:0;comment:难度" json:"difficulty,omitempty"`
	TimeLimit    float64           `gorm:"not null;default:1;comment:时间限制（s）" json:"time_limit,omitempty"`
	MemoryLimit  uint64            `gorm:"not null;default:131072;comment:内存限制（kb）" json:"memory_limit,omitempty"`
	Description  string            `gorm:"type:longtext;not null;comment:题面" json:"description,omitempty"`
	Input        string            `gorm:"type:longtext;not null;comment:输入说明" json:"input,omitempty"`
	Output       string            `gorm:"type:longtext;not null;comment:输出说明" json:"output,omitempty"`
	SampleInput  string            `gorm:"type:longtext;not null;comment:输入样例" json:"sample_input,omitempty"`
	SampleOutput string            `gorm:"type:longtext;not null;comment:输出样例" json:"sample_output,omitempty"`
	Hint         string            `gorm:"type:longtext;not null;comment:提示" json:"hint,omitempty"`
	Operation    Operation         `gorm:"not null;default:0;comment:操作" json:"operation,omitempty"`
	CreateTime   time.Time         `gorm:"not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time,omitempty"`
}

func (ProblemHistory) TableName() string {
	return "tbl_problem_history"
}

// 题目信息（题目+标签+评测点数据+题解）
type ProblemInfo struct {
	Problem   Problem    `json:"problem,omitempty"`
	Tags      []Tag      `json:"tags,omitempty"`
	Testcases []Testcase `json:"testcases,omitempty"`
	Solutions []Solution `json:"solution,omitempty"`
}

type CommonTestcaseInput struct {
	Rows []CommonTestcaseRow `json:"rows,omitempty"`
}

type CommonTestcaseRow struct {
	RowSizeId uint64                `json:"row_size_id,omitempty"`
	Values    []CommonTestcaseValue `json:"values,omitempty"`
}

type CommonTestcaseValue struct {
	ValueSizeId uint64  `json:"value_size_id,omitempty"`
	Type        string  `json:"type,omitempty"`
	Max         float64 `json:"max,omitempty"`
	Min         float64 `json:"min,omitempty"`
	MaxId       uint64  `json:"max_id,omitempty"`
	MinId       uint64  `json:"min_id,omitempty"`
}

func (c *CommonTestcaseInput) Unfold() utils.CommonTestcaseInput {
	rand.Seed(uint64(time.Now().UnixNano()))
	var input utils.CommonTestcaseInput
	var hsh []float64
	hsh = append(hsh, 0)
	for _, row := range c.Rows {
		if row.RowSizeId > 0 && row.RowSizeId < uint64(len(hsh)) {
			for i := 0; i < int(hsh[row.RowSizeId]); i++ {
				input.AppendRow(row.Unfold(&hsh))
			}
		} else {
			input.AppendRow(row.Unfold(&hsh))
		}
	}
	return input
}

func (c *CommonTestcaseRow) Unfold(hsh *[]float64) utils.CommonTestcaseRow {
	var row utils.CommonTestcaseRow
	for _, v := range c.Values {
		if v.ValueSizeId > 0 && v.ValueSizeId < uint64(len(*hsh)) {
			for i := 0; i < int((*hsh)[v.ValueSizeId]); i++ {
				row.AppendValue(v.Unfold(hsh))
			}
		} else {
			row.AppendValue(v.Unfold(hsh))
		}
	}
	return row
}

func (c *CommonTestcaseValue) Unfold(hsh *[]float64) utils.CommonTestcaseValue {
	if c.MaxId > 0 && c.MaxId < uint64(len(*hsh)) {
		c.Max = (*hsh)[c.MaxId]
	}
	if c.MinId > 0 && c.MinId < uint64(len(*hsh)) {
		c.Min = (*hsh)[c.MinId]
	}
	t := utils.GetValueType(c.Type)
	v := rand.Float64()*(c.Max-c.Min) + c.Min
	*hsh = append(*hsh, v)
	return utils.CommonTestcaseValue{
		Type:  t,
		Value: v,
	}
}