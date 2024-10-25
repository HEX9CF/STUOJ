package bootstrap

import (
	"STUOJ/conf"
	"log"
)

func InitConfig() {
	err := conf.InitConfig()
	if err != nil {
		log.Println("Init config failed!")
		panic(err)
	}
	log.Println("Init config success!")
}
