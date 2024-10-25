package bootstrap

import (
	"STUOJ/db"
	"STUOJ/judge"
	"STUOJ/utils"
	"log"
)

func InitJudge(chFin chan string) {
	var err error
	err = judge.InitJudge()
	if err != nil {
		log.Println("Init judge failed!")
	}

	err = InitJudgePrintInfo()
	if err != nil {
		log.Println("Init judge failed!")
	}

	err = InitJudgeLanguages()
	if err != nil {
		log.Println("Init judge failed!")
	}

	chFin <- "judge"
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

		_, err := db.InsertLanguage(language)
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
	if configtmp, err := utils.PrettyStruct(config); err != nil {
		log.Println("Struct formatting failed:", err)
		log.Println("Judge config info:", config)
	} else {
		log.Println("Judge config info:", configtmp)
	}

	system, err := judge.GetSystemInfo()
	if err != nil {
		return err
	}
	if systemtmp, err := utils.PrettyStruct(system); err != nil {
		log.Println("Struct formatting failed:", err)
		log.Println("Judge system info:", system)
	} else {
		log.Println("Judge system info:", systemtmp)
	}

	/*	statistics, err := judge.GetStatistics()
		if err != nil {
			return err
		}
		if statstmp, err := utils.PrettyStruct(statistics); err != nil {
			log.Println("Struct formatting failed:", err)
			log.Println("Judge statistics:", statistics)
		} else {
			log.Println("Judge statistics:", statstmp)
		}
	*/

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
		if workerstmp, err := utils.PrettyStruct(worker); err != nil {
			log.Println("Struct formatting failed:", err)
			log.Println(worker)
		} else {
			log.Println(workerstmp)
		}
	}

	/*	license, err := judge.GetLicense()
		if err != nil {
			return err
		}
		log.Println("Judge license:", license)
	*/

	/*	isolate, err := judge.GetIsolate()
		if err != nil {
			return err
		}
		log.Println("Judge isolate:", isolate)
	*/

	version, err := judge.GetVersion()
	if err != nil {
		return err
	}
	log.Println("Judge version:", version)

	return nil
}
