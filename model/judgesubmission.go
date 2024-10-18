package model

type JudgeSubmission struct {
	SourceCode string `json:"source_code"`
	LanguageId uint64 `json:"language_id"`
	Stdin string `json:"stdin"`
	ExpectedOutput string `json:"expected_output"`
	CPUTimeLimit float64 `json:"cpu_time_limit"`
	MemoryLimit uint64 `json:"memory_limit"`
}
