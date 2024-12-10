package bootstrap

import (
	"STUOJ/internal/conf"
	"log"
)

func initConfig() {
	err := conf.InitConfig()
	if err != nil {
		log.Println("Init config failed!")
		panic(err)
	}
	log.Println("Init config success.")
}
