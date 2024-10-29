package bootstrap

import (
	"STUOJ/yuki"
	"log"
)

func InitYuki() {
	err := yuki.InitYukiImage()
	if err != nil {
		log.Println("yuki-image init error:", err)
	}
}
