package bootstrap

import (
	"STUOJ/external/yuki"
	"log"
)

func InitYuki() {
	err := yuki.InitYukiImage()
	if err != nil {
		log.Println("yuki-image init error:", err)
	}
}
