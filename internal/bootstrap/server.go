package bootstrap

import (
	"STUOJ/server"
	"log"
)

func InitServer() {
	err := server.InitServer()
	if err != nil {
		log.Println("Init handler failed!")
		panic(err)
	}
	log.Println("Init handler success.")
}
