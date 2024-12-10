package model

// 题目
type NekoProblem struct {
	Title        string   `json:"title,omitempty"`
	Description  string   `json:"description,omitempty"`
	Input        string   `json:"input,omitempty"`
	Output       string   `json:"output,omitempty"`
	SampleInput  string   `json:"sample_input,omitempty"`
	SampleOutput string   `json:"sample_output,omitempty"`
	Hint         string   `json:"hint,omitempty"`
	Tags         []string `json:"tags,omitempty"`
}

// 题目说明
type NekoProblemInstruction struct {
	Title        string   `json:"title,omitempty" binding:"omitempty"`
	Description  string   `json:"description,omitempty" binding:"omitempty"`
	Input        string   `json:"input,omitempty" binding:"omitempty"`
	Output       string   `json:"output,omitempty" binding:"omitempty"`
	SampleInput  string   `json:"sample_input,omitempty" binding:"omitempty"`
	SampleOutput string   `json:"sample_output,omitempty" binding:"omitempty"`
	Hint         string   `json:"hint,omitempty" binding:"omitempty"`
	Tags         []string `json:"tags,omitempty" binding:"omitempty"`
	Solution     string   `json:"solution,omitempty" binding:"omitempty"`
}

// 测试用例
type NekoTestcase struct {
	TestInput         string `json:"test_input,omitempty"`
	TestOutput        string `json:"test_output,omitempty"`
	InputExplanation  string `json:"input_explanation,omitempty"`
	OutputExplanation string `json:"output_explanation,omitempty"`
}

// 测试用例说明
type NekoTestcaseInstruction struct {
	Title        string   `json:"title,omitempty" binding:"omitempty"`
	Description  string   `json:"description,omitempty" binding:"omitempty"`
	Input        string   `json:"input,omitempty" binding:"omitempty"`
	Output       string   `json:"output,omitempty" binding:"omitempty"`
	SampleInput  string   `json:"sample_input,omitempty" binding:"omitempty"`
	SampleOutput string   `json:"sample_output,omitempty" binding:"omitempty"`
	Hint         string   `json:"hint,omitempty" binding:"omitempty"`
	Tags         []string `json:"tags,omitempty" binding:"omitempty"`
	Solution     string   `json:"solution,omitempty" binding:"omitempty"`
}

// 题解
type NekoSolution struct {
	Language    string `json:"language,omitempty"`
	SourceCode  string `json:"source_code,omitempty"`
	Explanation string `json:"explanation,omitempty"`
}

// 题解说明
type NekoSolutionInstruction struct {
	Title        string   `json:"title,omitempty" binding:"omitempty"`
	Description  string   `json:"description,omitempty" binding:"omitempty"`
	Input        string   `json:"input,omitempty" binding:"omitempty"`
	Output       string   `json:"output,omitempty" binding:"omitempty"`
	SampleInput  string   `json:"sample_input,omitempty" binding:"omitempty"`
	SampleOutput string   `json:"sample_output,omitempty" binding:"omitempty"`
	Hint         string   `json:"hint,omitempty" binding:"omitempty"`
	Tags         []string `json:"tags,omitempty" binding:"omitempty"`
	Solution     string   `json:"solution,omitempty" binding:"omitempty"`
	Language     string   `json:"language,omitempty" binding:"omitempty"`
}

type NekoResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
