package judge

import (
	"net/http"
	"STUOJ/model"
	"encoding/json"
	"io/ioutil"
	"bytes"
)

func Submit(submission model.JudgeSubmission)(token string){
	data, err := json.Marshal(submission)
	if err != nil {
		return ""
	}
	url:=preUrl+"/submissions"
	req,_:=http.NewRequest("POST",url,bytes.NewReader(data))
	req.Header.Set("X-Auth-Token",config.Token)
	req.Header.Set("Content-Type","application/json")
	res,_:=http.DefaultClient.Do(req)
	defer res.Body.Close()
	body,_:=ioutil.ReadAll(res.Body)
	bodystr:=string(body)
	return bodystr
}