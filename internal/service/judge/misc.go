package judge

import (
	"STUOJ/external/judge0"
	"STUOJ/internal/model"
	"errors"
	"log"
)

func GetStatistics() (model.Judge0Statistics, error) {
	var err error
	var statistics model.Judge0Statistics

	// 获取评测机统计信息
	statistics.JudgeStatistics, err = judge0.GetStatistics()
	if err != nil {
		log.Println(err)
		return model.Judge0Statistics{}, errors.New("获取评测机统计信息失败")
	}

	// 获取评测机系统信息
	statistics.JudgeSystemInfo, err = judge0.GetSystemInfo()
	if err != nil {
		log.Println(err)
		return model.Judge0Statistics{}, errors.New("获取评测机系统信息失败")
	}

	return statistics, nil
}
