package bootstrap

import (
	"STUOJ/database"
	"log"
)

func InitDatabase() {
	err := database.InitDatabase()
	if err != nil {
		log.Println("Init database failed!")
		panic(err)
	}
	log.Println("Init database success.")
}
