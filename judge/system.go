package judge

import (
	"encoding/json"
	"STUOJ/model"
)

func GetLanguage() ([]model.Language, error) {
	bodystr,err:= httpInteraction("/languages","GET",nil)
	if err != nil {
		return nil,err
	}
	var languages []model.Language
	err=json.Unmarshal([]byte(bodystr), &languages)
	if err != nil {
		return nil, err
	}
	return languages, nil
}
