package record

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/model"
	"errors"
	"log"
)

// 查询所有提交记录（不返回源代码）
func SelectAll() ([]model.Record, error) {
	var records []model.Record

	// 获取提交信息
	submissions, err := dao.SelectAllSubmissions()
	if err != nil {
		log.Println(err)
		return nil, errors.New("获取提交信息失败")
	}

	// 不返回源代码
	for k := range submissions {
		submissions[k].SourceCode = ""
	}

	// 封装提交记录
	for _, s := range submissions {
		r := model.Record{
			Submission: s,
		}
		records = append(records, r)
	}

	return records, nil
}

// 根据提交ID查询提交记录
func SelectBySubmissionId(sid uint64) (model.Record, error) {
	// 获取提交信息
	s, err := dao.SelectSubmissionById(sid)
	if err != nil {
		log.Println(err)
		return model.Record{}, errors.New("获取提交信息失败")
	}

	// 获取评测结果
	judgements, err := dao.SelectJudgementsBySubmissionId(sid)
	if err != nil {
		log.Println(err)
		return model.Record{}, errors.New("获取评测结果失败")
	}

	// 封装提交记录
	r := model.Record{
		Submission: s,
		Judgements: judgements,
	}

	return r, nil
}

// 根据题目ID查询提交记录（不返回源代码）
func SelectByProblemId(problemId uint64) ([]model.Record, error) {
	var records []model.Record

	// 获取提交信息
	submissions, err := dao.SelectSubmissionsByProblemId(problemId)
	if err != nil {
		log.Println(err)
		return nil, errors.New("获取提交信息失败")
	}

	// 不返回源代码
	for k := range submissions {
		submissions[k].SourceCode = ""
	}

	// 封装提交记录
	for _, s := range submissions {
		r := model.Record{
			Submission: s,
		}
		records = append(records, r)
	}

	return records, nil
}

// 根据用户ID查询提交记录（不返回源代码）
func SelectByUserId(userId uint64) ([]model.Record, error) {
	var records []model.Record

	// 获取提交信息
	submissions, err := dao.SelectSubmissionsByUserId(userId)
	if err != nil {
		log.Println(err)
		return nil, errors.New("获取提交信息失败")
	}

	// 不返回源代码
	for k := range submissions {
		submissions[k].SourceCode = ""
	}

	// 封装提交记录
	for _, s := range submissions {
		r := model.Record{
			Submission: s,
		}
		records = append(records, r)
	}

	return records, nil
}
