package bootstrap

import (
	"STUOJ/db"
	"STUOJ/judge"
	"log"
)

func InitJudge() error {
	var err error
	err = judge.InitJudge()
	if err != nil {
		return err
	}

	err = InitJudgePrintInfo()
	if err != nil {
		return err
	}

	err = InitJudgeLanguages()
	if err != nil {
		return err
	}

	return nil
}

// 初始化评测机语言
func InitJudgeLanguages() error {
	// 读取评测机语言列表
	languages, err := judge.GetLanguage()
	if err != nil {
		return err
	}

	log.Println("Judge support languages:")
	for k, v := range languages {
		log.Println(k, v)
	}

	// 清空数据库语言表
	err = db.DeleteAllLanguages()
	if err != nil {
		log.Println(err)
	}

	// 插入数据库语言表
	for _, language := range languages {
		//log.Println(v)

		// 初始化对象
		// language := model.Language{
		// 	Id:   uint64(v["id"].(float64)),
		// 	Name: v["name"].(string),
		// }

		err := db.InsertLanguage(language)
		if err != nil {
			return err
		}
	}

	return nil
}

// 打印评测机信息
func InitJudgePrintInfo() error {
	config, err := judge.GetConfigInfo()
	if err != nil {
		return err
	}
	log.Println("Judge config info:", config)

	system, err := judge.GetSystemInfo()
	if err != nil {
		return err
	}
	log.Println("Judge system info:", system)

	statistics, err := judge.GetStatistics()
	if err != nil {
		return err
	}
	log.Println("Judge statistics:", statistics)

	/*	about, err := judge.GetAbout()
		if err != nil {
			return err
		}
		log.Println("Judge about:", about)
	*/
	workers, err := judge.GetWorkers()
	if err != nil {
		return err
	}
	log.Println("Judge workers:")
	for _, worker := range workers {
		log.Println(worker)
	}

	/*	license, err := judge.GetLicense()
		if err != nil {
			return err
		}
		log.Println("Judge license:", license)
	*/
	isolate, err := judge.GetIsolate()
	if err != nil {
		return err
	}
	log.Println("Judge isolate:", isolate)

	version, err := judge.GetVersion()
	if err != nil {
		return err
	}
	log.Println("Judge version:", version)

	return nil
}
