package judge

import (
	"STUOJ/external/judge0"
	"STUOJ/internal/dao"
	"STUOJ/internal/model"
	"errors"
	"log"
)

func GetStatistics() (model.Judge0Statistics, error) {
	var err error
	var stats model.Judge0Statistics

	// 统计语言数量
	stats.LanguageCount, err = dao.CountLanguages()
	if err != nil {
		log.Println(err)
		return model.Judge0Statistics{}, errors.New("统计语言数量失败")
	}

	// 获取评测机统计信息
	stats.JudgeStatistics, err = judge0.GetStatistics()
	if err != nil {
		log.Println(err)
		return model.Judge0Statistics{}, errors.New("获取评测机统计信息失败")
	}

	// 获取评测机系统信息
	stats.JudgeSystemInfo, err = judge0.GetSystemInfo()
	if err != nil {
		log.Println(err)
		return model.Judge0Statistics{}, errors.New("获取评测机系统信息失败")
	}

	return stats, nil
}
