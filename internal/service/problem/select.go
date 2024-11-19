package problem

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"errors"
	"log"
)

// 根据ID查询题目数据
func SelectById(id uint64) (model.ProblemData, error) {
	// 获取题目信息
	p, err := dao.SelectProblemById(id)
	if err != nil {
		return model.ProblemData{}, errors.New("获取题目信息失败")
	}

	// 获取题目标签
	tags, err := dao.SelectTagsByProblemId(id)
	if err != nil {
		return model.ProblemData{}, errors.New("获取题目标签失败")
	}

	// 获取评测点数据
	testcases, err := dao.SelectTestcasesByProblemId(id)
	if err != nil {
		return model.ProblemData{}, errors.New("获取评测点数据失败")
	}

	// 获取题解数据
	solutions, err := dao.SelectSolutionsByProblemId(id)
	if err != nil {
		return model.ProblemData{}, errors.New("获取题解数据失败")
	}

	// 不返回源代码
	for i := range solutions {
		solutions[i].SourceCode = ""
	}

	// 获取题目历史记录
	histories, err := dao.SelectHistoriesByProblemId(id)
	if err != nil {
		return model.ProblemData{}, errors.New("获取题目历史记录失败")
	}

	// 封装题目数据
	pd := model.ProblemData{
		Problem:   p,
		Tags:      tags,
		Testcases: testcases,
		Solutions: solutions,
		Histories: histories,
	}

	return pd, nil
}

// 查询所有题目数据
func SelectAll() ([]model.ProblemData, error) {
	problems, err := dao.SelectAllProblems()
	if err != nil {
		log.Println(err)
		return nil, errors.New("获取题目信息失败")
	}

	pds := wrapProblemDatas(problems)

	return pds, nil
}

// 根据状态查询题目数据
func SelectByStatus(s entity.ProblemStatus) ([]model.ProblemData, error) {
	problems, err := dao.SelectProblemsByStatus(s)
	if err != nil {
		log.Println(err)
		return nil, errors.New("获取题目信息失败")
	}

	pds := wrapProblemDatas(problems)

	return pds, nil
}

// 查询公开题目数据
func SelectPublic() ([]model.ProblemData, error) {
	problems, err := dao.SelectProblemsByStatus(entity.ProblemStatusPublic)
	if err != nil {
		log.Println(err)
		return nil, errors.New("获取题目信息失败")
	}

	pds := wrapProblemDatas(problems)

	return pds, nil
}

// 根据状态和标签查询题目
func SelectPublicByTagId(tid uint64) ([]model.ProblemData, error) {
	problems, err := dao.SelectProblemsByTagIdAndStatus(tid, entity.ProblemStatusPublic)
	if err != nil {
		log.Println(err)
		return nil, errors.New("获取题目信息失败")
	}

	pds := wrapProblemDatas(problems)

	return pds, nil
}

// 根据状态和难度查询题目
func SelectPublicByDifficulty(d entity.Difficulty) ([]model.ProblemData, error) {
	problems, err := dao.SelectProblemsByDifficultyAndStatus(d, entity.ProblemStatusPublic)
	if err != nil {
		return nil, errors.New("获取题目信息失败")
	}

	pds := wrapProblemDatas(problems)

	return pds, nil
}

// 根据状态查询并根据标题模糊查询公开题目
func SelectPublicLikeTitle(title string) ([]model.ProblemData, error) {
	problems, err := dao.SelectProblemsLikeTitleByStatus(title, entity.ProblemStatusPublic)
	if err != nil {
		log.Println(err)
		return nil, errors.New("获取题目信息失败")
	}

	pds := wrapProblemDatas(problems)

	return pds, nil
}

// 根据题目ID查询公开题目数据
func SelectPublicByProblemId(pid uint64) (model.ProblemData, error) {
	p, err := dao.SelectProblemByIdAndStatus(pid, entity.ProblemStatusPublic)
	if err != nil {
		log.Println(err)
		return model.ProblemData{}, errors.New("获取题目信息失败")
	}

	// 获取题目标签
	tags, err := dao.SelectTagsByProblemId(pid)
	if err != nil {
		log.Println(err)
		return model.ProblemData{}, errors.New("获取题目标签失败")
	}

	// 初始化题目信息
	pd := model.ProblemData{
		Problem: p,
		Tags:    tags,
	}

	return pd, nil
}

func hideProblemContent(problems []entity.Problem) {
	for i := range problems {
		problems[i].Description = ""
		problems[i].Input = ""
		problems[i].Output = ""
		problems[i].SampleInput = ""
		problems[i].SampleOutput = ""
		problems[i].Hint = ""
	}
}

// 封装题目数据
func wrapProblemDatas(problems []entity.Problem) []model.ProblemData {
	var pds []model.ProblemData

	hideProblemContent(problems)

	for _, p := range problems {
		pd := model.ProblemData{
			Problem: p,
		}

		pds = append(pds, pd)
	}

	return pds
}
