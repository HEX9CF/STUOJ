package fps

import (
	"STUOJ/internal/db"
	model2 "STUOJ/internal/model"
)

func Parse(fps model2.FPS) []model2.ProblemInfo {
	problems := make([]model2.ProblemInfo, 0)
	for _, item := range fps.Items {
		problems = append(problems, ParseItem(item))
	}
	return problems
}

func ParseItem(item model2.Item) model2.ProblemInfo {
	problem := item.ToProblem()
	testcases := item.GetTestCase()
	var solutions []model2.Solution
	for _, solution := range item.Solution {
		var languageId uint64
		language, err := db.SelectLanguageLikeName(solution.Language)
		if err != nil {
			languageId = 0
		} else {
			languageId = language.Id
		}
		solutions = append(solutions, model2.Solution{LanguageId: languageId, SourceCode: solution.Code})
	}
	return model2.ProblemInfo{
		Problem:   problem,
		Testcases: testcases,
		Solutions: solutions,
	}
}
