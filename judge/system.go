package judge

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"STUOJ/model"
)

func GetLanguage() ([]model.Language, error) {
	url := preUrl + "/languages"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("X-Auth-Token", config.Token)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	bodystr := string(body)
	var languages []model.Language
	err:=json.Unmarshal([]byte(bodystr), &languages)
	if err != nil {
		return nil, err
	}
	return languages, nil
}
