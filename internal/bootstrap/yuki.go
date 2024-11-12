package bootstrap

import (
	"STUOJ/external/yuki"
	"log"
)

func initYuki() {
	err := yuki.InitYukiImage()
	if err != nil {
		log.Println("yuki-image init error:", err)
	}
}
