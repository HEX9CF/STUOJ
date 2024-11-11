package record

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/model"
)

// 根据提交ID查询评测结果
func SelectJudgementsBySubmissionId(sid uint64) ([]model.Judgement, error) {
	judgements, err := dao.SelectJudgementsBySubmissionId(sid)
	if err != nil {
		return nil, err
	}

	return judgements, nil
}

// 查询所有提交记录（不返回源代码）
func SelectAllSubmissions() ([]model.Submission, error) {
	submissions, err := dao.SelectAllSubmissions()
	if err != nil {
		return nil, err
	}

	// 不返回源代码
	for k := range submissions {
		submissions[k].SourceCode = ""
	}

	return submissions, nil
}

// 根据ID查询提交记录
func SelectSubmissionById(id uint64) (model.Submission, error) {
	s, err := dao.SelectSubmissionById(id)
	if err != nil {
		return model.Submission{}, err
	}

	return s, nil
}

// 根据用户ID查询提交记录（不返回源代码）
func SelectSubmissionsByUserId(userId uint64) ([]model.Submission, error) {
	submissions, err := dao.SelectSubmissionsByUserId(userId)
	if err != nil {
		return nil, err
	}

	// 不返回源代码
	for k := range submissions {
		submissions[k].SourceCode = ""
	}

	return submissions, nil
}

// 根据题目ID查询提交记录（不返回源代码）
func SelectSubmissionsByProblemId(problemId uint64) ([]model.Submission, error) {
	submissions, err := dao.SelectSubmissionsByProblemId(problemId)
	if err != nil {
		return nil, err
	}

	// 不返回源代码
	for k := range submissions {
		submissions[k].SourceCode = ""
	}

	return submissions, nil
}
