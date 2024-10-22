package bootstrap

import (
	"STUOJ/conf"
	"log"
)

func Init() {
	var err error

	err = conf.InitConfig()
	if err != nil {
		log.Println("Init config failed!")
		panic(err)
	}

	err = InitDatabase()
	if err != nil {
		log.Println("Init database failed!")
		panic(err)
	}

	err = InitJudge()
	if err != nil {
		log.Println("Init judge failed!")
		panic(err)
	}

	err = InitLskypro()
	if err != nil {
		log.Println("Init lskypro failed!")
		panic(err)
	}

	err = InitHandlers()
	if err != nil {
		log.Println("Init handlers failed!")
		panic(err)
	}

}
