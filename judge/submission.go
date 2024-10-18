package judge

import (
	"STUOJ/model"
	"encoding/json"
	"bytes"
)

func Submit(submission model.JudgeSubmission)(string,error){
	data, err := json.Marshal(submission)
	if err != nil {
		return "",err
	}
	bodystr,err:= httpInteraction("/submissions","POST",bytes.NewReader(data))
	return bodystr,nil
}

func Query(token string)(model.JudgeResult,error){
	bodystr,err:= httpInteraction("/submissions"+"/"+token,"GET",nil)
	if err != nil {
		return model.JudgeResult{},err
	}
	var result model.JudgeResult
	err=json.Unmarshal([]byte(bodystr),&result)
	if err != nil {
		return model.JudgeResult{},err
	}
	return result,nil
}