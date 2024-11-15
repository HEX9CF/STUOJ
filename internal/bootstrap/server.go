package bootstrap

import (
	"STUOJ/server"
	"log"
)

func initServer() {
	err := server.InitServer()
	if err != nil {
		log.Println("Init server failed!")
		panic(err)
	}
	log.Println("Init server success.")
}
