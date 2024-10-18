package judge

import (
	"net/http"
	"STUOJ/model"
	"encoding/json"
	"io/ioutil"
	"bytes"
)

func Submit(submission model.JudgeSubmission)(string,error){
	data, err := json.Marshal(submission)
	if err != nil {
		return "",err
	}
	url:=preUrl+"/submissions"
	req,err:=http.NewRequest("POST",url,bytes.NewReader(data))
	if err != nil {
		return "",err
	}
	req.Header.Set("X-Auth-Token",config.Token)
	req.Header.Set("Content-Type","application/json")
	res,err:=http.DefaultClient.Do(req)
	if err != nil {
		return "",err
	}
	defer res.Body.Close()
	body,err:=ioutil.ReadAll(res.Body)
	if err != nil {
		return "",err
	}
	bodystr:=string(body)
	return bodystr,nil
}

func Query(token string)(model.JudgeResult,error){
	url:=preUrl+"/submissions"+"/"+token
	req,err:=http.NewRequest("GET",url,nil)
	if err != nil {
		return model.JudgeResult{},err
	}
	req.Header.Set("X-Auth-Token",config.Token)
	res,err:=http.DefaultClient.Do(req)
	if err != nil {
		return model.JudgeResult{},err
	}
	defer res.Body.Close()
	body,err:=ioutil.ReadAll(res.Body)
	if err != nil {
		return model.JudgeResult{},err
	}
	bodystr:=string(body)
	var result model.JudgeResult
	err=json.Unmarshal([]byte(bodystr),&result)
	if err != nil {
		return model.JudgeResult{},err
	}
	return result,nil
}