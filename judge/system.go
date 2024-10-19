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

func GetConfigInfo()(model.JudgeConfigInfo,error){
	bodystr,err:= httpInteraction("/config_info","GET",nil)
	if err != nil {
		return model.JudgeConfigInfo{},err
	}
	var config model.JudgeConfigInfo
	err=json.Unmarshal([]byte(bodystr), &config)
	if err != nil {
		return model.JudgeConfigInfo{}, err
	}
	return config, nil
}

func GetSystemInfo()(model.JudgeSystemInfo,error){
	bodystr,err:= httpInteraction("/system_info","GET",nil)
	if err != nil {
		return model.JudgeSystemInfo{},err
	}
	var system model.JudgeSystemInfo
	err=json.Unmarshal([]byte(bodystr), &system)
	if err != nil {
		return model.JudgeSystemInfo{}, err
	}
	return system, nil
}

func GetStatistics()(model.JudgeStatistics,error){
	bodystr,err:= httpInteraction("/statistics","GET",nil)
	if err != nil {
		return model.JudgeStatistics{},err
	}
	var statistics model.JudgeStatistics
	err=json.Unmarshal([]byte(bodystr), &statistics)
	if err != nil {
		return model.JudgeStatistics{}, err
	}
	return statistics, nil
}

func GetAbout()(model.JudgeAbout,error){
	bodystr,err:= httpInteraction("/about","GET",nil)
	if err != nil {
		return model.JudgeAbout{},err
	}
	var about model.JudgeAbout
	err=json.Unmarshal([]byte(bodystr), &about)
	if err != nil {
		return model.JudgeAbout{}, err
	}
	return about, nil
}

func GetWorkers()([]model.JudgeWorker,error){
	bodystr,err:= httpInteraction("/workers","GET",nil)
	if err != nil {
		return nil,err
	}
	var workers []model.JudgeWorker
	err=json.Unmarshal([]byte(bodystr), &workers)
	if err != nil {
		return nil, err
	}
	return workers, nil
}