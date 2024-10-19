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

func GetConfigInfo()(model.ConfigInfo,error){
	bodystr,err:= httpInteraction("/config_info","GET",nil)
	if err != nil {
		return model.ConfigInfo{},err
	}
	var config model.ConfigInfo
	err=json.Unmarshal([]byte(bodystr), &config)
	if err != nil {
		return model.ConfigInfo{}, err
	}
	return config, nil
}