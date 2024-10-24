package judge

import (
	"STUOJ/model"
	"bytes"
	"encoding/json"
	"strconv"
)

func Submit(submission model.JudgeSubmission) (model.JudgeSubmissionResult, error) {
	data, err := json.Marshal(submission)
	if err != nil {
		return model.JudgeSubmissionResult{}, err
	}
	bodystr, err := httpInteraction("/submissions", "POST", bytes.NewReader(data))
	if err != nil {
		return model.JudgeSubmissionResult{}, err
	}
	var result model.JudgeSubmissionResult
	err = json.Unmarshal([]byte(bodystr), &result)
	if err != nil {
		return model.JudgeSubmissionResult{}, err
	}

	return result, nil
}

func QueryResult(token string) (model.JudgeSubmissionResult, error) {
	bodystr, err := httpInteraction("/submissions"+"/"+token, "GET", nil)
	if err != nil {
		return model.JudgeSubmissionResult{}, err
	}
	var result model.JudgeSubmissionResult
	err = json.Unmarshal([]byte(bodystr), &result)
	if err != nil {
		return model.JudgeSubmissionResult{}, err
	}
	return result, nil
}

func QueryResults(page uint64, per_page uint64) (model.JudgeSubmissionResults, error) {
	bodystr, err := httpInteraction("/submissions"+"/?page="+strconv.FormatUint(page, 10)+"&per_page="+strconv.FormatUint(per_page, 10), "GET", nil)
	if err != nil {
		return model.JudgeSubmissionResults{}, err
	}
	var results model.JudgeSubmissionResults
	err = json.Unmarshal([]byte(bodystr), &results)
	if err != nil {
		return model.JudgeSubmissionResults{}, err
	}
	return results, nil
}
