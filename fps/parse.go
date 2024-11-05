package fps

import "STUOJ/model"

func Parse(fps model.FPS) ([]model.Problem, error) {
	var problems []model.Problem
	for _, fp := range fps.Items {
		problems = append(problems, fp.ToProblem())
	}
	return problems, nil
}

func ParseItem(item model.Item) (model.Problem, []model.Testcase) {
	problem := item.ToProblem()
	testcases := item.GetTestCase()
	return problem, testcases
}
