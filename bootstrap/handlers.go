package bootstrap

import (
	"STUOJ/handlers"
	"log"
)

func InitHandlers() {
	err := handlers.InitHandlers()
	if err != nil {
		log.Println("Init handlers failed!")
		panic(err)
	}
	log.Println("Init handlers success.")
}
