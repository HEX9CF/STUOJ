package fps

import (
	"STUOJ/db"
	"STUOJ/model"
)

func Parse(fps model.FPS) ([]model.Problem, error) {
	var problems []model.Problem
	for _, fp := range fps.Items {
		problems = append(problems, fp.ToProblem())
	}
	return problems, nil
}

func ParseItem(item model.Item) (model.Problem, []model.Testcase, []model.Solution) {
	problem := item.ToProblem()
	testcases := item.GetTestCase()
	var solutions []model.Solution
	for _, solution := range item.Solution {
		var languageId uint64
		language, err := db.SelectLanguageLikeName(solution.Language)
		if err != nil {
			languageId = 0
		} else {
			languageId = language.Id
		}
		solutions = append(solutions, model.Solution{LanguageId: languageId, SourceCode: solution.Code})
	}
	return problem, testcases, solutions
}
