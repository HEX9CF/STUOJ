package fps

import (
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"STUOJ/internal/service/language"
)

func Parse(fps model.FPS) []model.ProblemData {
	problems := make([]model.ProblemData, 0)
	for _, item := range fps.Items {
		problems = append(problems, ParseItem(item))
	}
	return problems
}

func ParseItem(item model.Item) model.ProblemData {
	problem := item.ToProblem()
	testcases := item.GetTestCase()
	var solutions []entity.Solution
	for _, solution := range item.Solution {
		var languageId uint64
		l, err := language.SelectLikeName(solution.Language)
		if err != nil {
			languageId = 0
		} else {
			languageId = l.Id
		}
		solutions = append(solutions, entity.Solution{LanguageId: languageId, SourceCode: solution.Code})
	}
	return model.ProblemData{
		Problem:   problem,
		Testcases: testcases,
		Solutions: solutions,
	}
}
