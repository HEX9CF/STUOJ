package bootstrap

import (
	"STUOJ/external/yuki"
	"log"
)

func initYuki() {
	err := yuki.InitYukiImage()
	if err != nil {
		log.Println(err)
		log.Println("Init yuki-image error!")
		return
	}
	log.Println("Init yuki-image success.")
}
