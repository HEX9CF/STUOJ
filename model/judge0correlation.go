package model

type JudgeSubmission struct {
	SourceCode string `json:"source_code"`
	LanguageId uint64 `json:"language_id"`
	Stdin string `json:"stdin"`
	ExpectedOutput string `json:"expected_output"`
	CPUTimeLimit float64 `json:"cpu_time_limit"`
	MemoryLimit uint64 `json:"memory_limit"`
}

type JudgeResult struct{
	Stdout string `json:"stdout"`
	Time string `json:"time"`
	Memory float64 `json:"memory"`
	Stderr string `json:"stderr"`
	Token string `json:"token"`
	Message string `json:"message"`
	Status JudgeStatus `json:"status"`
}

type JudgeStatus struct{
	Id uint64 `json:"id"`
	Description string `json:"description"`
}