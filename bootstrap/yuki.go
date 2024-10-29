package bootstrap

import (
	"STUOJ/yuki"
	"log"
)

func InitYuki(chFin chan string) {
	err := yuki.InitYukiImage()
	if err != nil {
		log.Println("yuki-image init error:", err)
	}
	chFin <- "yuki-image init success!"
}
