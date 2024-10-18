package judge

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"STUOJ/model"
)

func GetLanguage() ([]model.Language, error) {
	url := preUrl + "/languages"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil,err
	}
	req.Header.Set("X-Auth-Token", config.Token)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil,err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil,err
	}
	bodystr := string(body)
	var languages []model.Language
	err=json.Unmarshal([]byte(bodystr), &languages)
	if err != nil {
		return nil, err
	}
	return languages, nil
}
