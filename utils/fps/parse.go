package fps

import (
	"STUOJ/internal/model"
	"STUOJ/internal/service/language"
)

func Parse(fps model.FPS) []model.ProblemInfo {
	problems := make([]model.ProblemInfo, 0)
	for _, item := range fps.Items {
		problems = append(problems, ParseItem(item))
	}
	return problems
}

func ParseItem(item model.Item) model.ProblemInfo {
	problem := item.ToProblem()
	testcases := item.GetTestCase()
	var solutions []model.Solution
	for _, solution := range item.Solution {
		var languageId uint64
		language, err := language.SelectLikeName(solution.Language)
		if err != nil {
			languageId = 0
		} else {
			languageId = language.Id
		}
		solutions = append(solutions, model.Solution{LanguageId: languageId, SourceCode: solution.Code})
	}
	return model.ProblemInfo{
		Problem:   problem,
		Testcases: testcases,
		Solutions: solutions,
	}
}
