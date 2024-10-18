package bootstrap

import (
	"STUOJ/db"
	"STUOJ/judge"
	// "STUOJ/model"
	"log"
)

func InitJudge() {
	judge.InitJudge()
	InitJudgeLanguages()
}

func InitJudgeLanguages() {
	// 读取评测机语言列表
	languages, err := judge.GetLanguage()
	if err != nil {
		log.Println(err)
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
	for _,language := range languages {
		//log.Println(v)

		// 初始化对象
		// language := model.Language{
		// 	Id:   uint64(v["id"].(float64)),
		// 	Name: v["name"].(string),
		// }

		err := db.InsertLanguage(language)
		if err != nil {
			log.Println(err)
		}
	}
}
