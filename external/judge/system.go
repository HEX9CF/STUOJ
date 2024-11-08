package judge

import (
	model2 "STUOJ/internal/model"
	"encoding/json"
)

func GetLanguage() ([]model2.Language, error) {
	bodystr, err := httpInteraction("/languages", "GET", nil)
	if err != nil {
		return nil, err
	}
	var languages []model2.Language
	err = json.Unmarshal([]byte(bodystr), &languages)
	if err != nil {
		return nil, err
	}
	return languages, nil
}

func GetConfigInfo() (model2.JudgeConfigInfo, error) {
	bodystr, err := httpInteraction("/config_info", "GET", nil)
	if err != nil {
		return model2.JudgeConfigInfo{}, err
	}
	var config model2.JudgeConfigInfo
	err = json.Unmarshal([]byte(bodystr), &config)
	if err != nil {
		return model2.JudgeConfigInfo{}, err
	}
	return config, nil
}

func GetSystemInfo() (model2.JudgeSystemInfo, error) {
	bodystr, err := httpInteraction("/system_info", "GET", nil)
	if err != nil {
		return model2.JudgeSystemInfo{}, err
	}
	var system model2.JudgeSystemInfo
	err = json.Unmarshal([]byte(bodystr), &system)
	if err != nil {
		return model2.JudgeSystemInfo{}, err
	}
	return system, nil
}

func GetStatistics() (model2.JudgeStatistics, error) {
	bodystr, err := httpInteraction("/statistics", "GET", nil)
	if err != nil {
		return model2.JudgeStatistics{}, err
	}
	var statistics model2.JudgeStatistics
	err = json.Unmarshal([]byte(bodystr), &statistics)
	if err != nil {
		return model2.JudgeStatistics{}, err
	}
	return statistics, nil
}

func GetAbout() (model2.JudgeAbout, error) {
	bodystr, err := httpInteraction("/about", "GET", nil)
	if err != nil {
		return model2.JudgeAbout{}, err
	}
	var about model2.JudgeAbout
	err = json.Unmarshal([]byte(bodystr), &about)
	if err != nil {
		return model2.JudgeAbout{}, err
	}
	return about, nil
}

func GetWorkers() ([]model2.JudgeWorker, error) {
	bodystr, err := httpInteraction("/workers", "GET", nil)
	if err != nil {
		return nil, err
	}
	var workers []model2.JudgeWorker
	err = json.Unmarshal([]byte(bodystr), &workers)
	if err != nil {
		return nil, err
	}
	return workers, nil
}

func GetLicense() (string, error) {
	bodystr, err := httpInteraction("/license", "GET", nil)
	if err != nil {
		return "", err
	}
	return bodystr, nil
}

func GetIsolate() (string, error) {
	bodystr, err := httpInteraction("/isolate", "GET", nil)
	if err != nil {
		return "", err
	}
	return bodystr, nil
}

func GetVersion() (string, error) {
	bodystr, err := httpInteraction("/version", "GET", nil)
	if err != nil {
		return "", err
	}
	return bodystr, nil
}
