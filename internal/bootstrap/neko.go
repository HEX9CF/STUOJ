package bootstrap

import (
	"STUOJ/external/neko"
	"log"
)

func initNeko() {
	err := neko.InitNekoAcm()
	if err != nil {
		log.Println(err)
		log.Println("Init NekoACM error!")
		return
	}
	log.Println("Init NekoACM success.")
}
