package record

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"errors"
	"log"
)

// 查询所有提交记录（不返回源代码）
func SelectAll(userId uint64, hideCode ...bool) ([]entity.Submission, error) {
	// 获取提交信息
	submissions, err := dao.SelectAllSubmissions()
	if err != nil {
		log.Println(err)
		return nil, errors.New("获取提交信息失败")
	}
	if len(hideCode) == 0 || hideCode[0] { // 隐藏源代码
		hideSubmissionSourceCode(userId, submissions)
	}

	return submissions, nil
}

// 根据提交ID查询提交记录
func SelectBySubmissionId(userId uint64, sid uint64, hideCode ...bool) (model.Record, error) {
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

	if (len(hideCode) == 0 || hideCode[0]) && userId != s.UserId { // 隐藏源代码
		s.SourceCode = ""
	}

	// 封装提交记录
	r := model.Record{
		Submission: s,
		Judgements: judgements,
	}

	return r, nil
}

// 根据题目ID查询提交记录（不返回源代码）
func SelectByProblemId(userId uint64, problemId uint64, hideCode ...bool) ([]entity.Submission, error) {
	// 获取提交信息
	submissions, err := dao.SelectSubmissionsByProblemId(problemId)
	if err != nil {
		log.Println(err)
		return nil, errors.New("获取提交信息失败")
	}

	if len(hideCode) == 0 || hideCode[0] { // 隐藏源代码
		hideSubmissionSourceCode(userId, submissions)
	}

	return submissions, nil
}

// 根据用户ID查询提交记录（不返回源代码）
func SelectByUserId(userId uint64, hideCode ...bool) ([]entity.Submission, error) {
	// 获取提交信息
	submissions, err := dao.SelectSubmissionsByUserId(userId)
	if err != nil {
		log.Println(err)
		return nil, errors.New("获取提交信息失败")
	}

	if len(hideCode) == 0 || hideCode[0] { // 隐藏源代码
		hideSubmissionSourceCode(0, submissions)
	}

	return submissions, nil
}

// 隐藏源代码
func hideSubmissionSourceCode(userId uint64, submissions []entity.Submission) {
	for i := range submissions {
		if submissions[i].UserId != userId {
			submissions[i].SourceCode = ""
		}
	}
}

// 封装提交记录
// func wrapRecords(submissions []entity.Submission) []model.Record {
// 	var records []model.Record

// 	hideSubmissionSourceCode(submissions)

// 	for _, s := range submissions {
// 		r := model.Record{
// 			Submission: s,
// 		}
// 		records = append(records, r)
// 	}

// 	return records
// }
