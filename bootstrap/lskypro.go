package bootstrap

import (
	"STUOJ/lskypro"
	"log"
)

func InitLskypro(chFin chan string) {
	err := lskypro.InitLskypro()
	if err != nil {
		log.Println("Init lskypro failed!")
	}

	chFin <- "lskypro"
}
