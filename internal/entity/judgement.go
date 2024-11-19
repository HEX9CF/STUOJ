package entity

// 评测状态：0 Pend, 1 In Queue, 2 Proc, 3 AC, 4 WA, 5 TLE, 6 CE, 7 RE(SIGSEGV), 8 RE(SIGXFSZ), 9 RE(SIGFPE), 10 RE(SIGABRT), 11 RE(NZEC), 12 RE(Other), 13 IE, 14 EFE
type JudgeStatus uint64

const (
	JudgeStatusPend      JudgeStatus = 0
	JudgeStatusIQ        JudgeStatus = 1
	JudgeStatusProc      JudgeStatus = 2
	JudgeStatusAC        JudgeStatus = 3
	JudgeStatusWA        JudgeStatus = 4
	JudgeStatusTLE       JudgeStatus = 5
	JudgeStatusCE        JudgeStatus = 6
	JudgeStatusRESIGSEGV JudgeStatus = 7
	JudgeStatusRESIGXFSZ JudgeStatus = 8
	JudgeStatusRESIGFPE  JudgeStatus = 9
	JudgeStatusRESIGABRT JudgeStatus = 10
	JudgeStatusRENZEC    JudgeStatus = 11
	JudgeStatusREOther   JudgeStatus = 12
	JudgeStatusIE        JudgeStatus = 13
	JudgeStatusEFE       JudgeStatus = 14
)

func (s JudgeStatus) String() string {
	switch s {
	case JudgeStatusPend:
		return "Pending"
	case JudgeStatusIQ:
		return "In Queue"
	case JudgeStatusProc:
		return "Processing"
	case JudgeStatusAC:
		return "Accepted"
	case JudgeStatusWA:
		return "Wrong Answer"
	case JudgeStatusTLE:
		return "Time Limit Exceeded"
	case JudgeStatusCE:
		return "Compilation Error"
	case JudgeStatusRESIGSEGV:
		return "Runtime Error (SIGSEGV)"
	case JudgeStatusRESIGXFSZ:
		return "Runtime Error (SIGXFSZ)"
	case JudgeStatusRESIGFPE:
		return "Runtime Error (SIGFPE)"
	case JudgeStatusRESIGABRT:
		return "Runtime Error (SIGABRT)"
	case JudgeStatusRENZEC:
		return "Runtime Error (NZEC)"
	case JudgeStatusREOther:
		return "Runtime Error (Other)"
	case JudgeStatusIE:
		return "Internal Error"
	case JudgeStatusEFE:
		return "Exec Format Error"
	default:
		return "Unknown"
	}
}

// 评测点结果
type Judgement struct {
	Id            uint64      `gorm:"primaryKey;autoIncrement;comment:评测点结果ID" json:"id,omitempty"`
	SubmissionId  uint64      `gorm:"not null;default:0;comment:提交记录ID" json:"submission_id,omitempty"`
	TestcaseId    uint64      `gorm:"not null;default:0;comment:评测点数据ID" json:"testcase_id,omitempty"`
	Time          float64     `gorm:"not null;default:0;comment:运行耗时（s）" json:"time,omitempty"`
	Memory        uint64      `gorm:"not null;default:0;comment:内存（kb）" json:"memory,omitempty"`
	Stdout        string      `gorm:"type:longtext;not null;comment:标准输出" json:"stdout,omitempty"`
	Stderr        string      `gorm:"type:longtext;not null;comment:标准错误输出" json:"stderr,omitempty"`
	CompileOutput string      `gorm:"type:longtext;not null;comment:编译输出" json:"compile_output,omitempty"`
	Message       string      `gorm:"type:longtext;not null;comment:信息" json:"message,omitempty"`
	Status        JudgeStatus `gorm:"not null;default:0;comment:状态" json:"status,omitempty"`
}

func (Judgement) TableName() string {
	return "tbl_judgement"
}
