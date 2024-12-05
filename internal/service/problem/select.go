package problem

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"errors"
)

// 根据ID查询题目数据
func SelectById(id uint64, admin ...bool) (model.ProblemData, error) {
	condition := dao.ProblemWhere{}
	condition.Id.Set(id)
	if len(admin) == 0 || !admin[0] {
		condition.Status.Set(entity.ProblemStatusPublic)
	}
	// 获取题目信息
	p, err := dao.SelectProblem(condition, 1, 10)
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

	// 封装题目数据
	pd := model.ProblemData{
		Problem:   p[0],
		Tags:      tags,
		Testcases: testcases,
		Solutions: solutions,
	}

	return pd, nil
}

func SelectProblem(condition dao.ProblemWhere, page uint64, size uint64) ([]model.ProblemData, error) {
	problems, err := dao.SelectProblem(condition, page, size)
	if err != nil {
		return []model.ProblemData{}, errors.New("获取题目信息失败")
	}

	pds := wrapProblemDatas(problems)

	return pds, nil
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
