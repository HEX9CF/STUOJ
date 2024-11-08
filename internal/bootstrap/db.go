package bootstrap

import (
	"STUOJ/internal/db"
	"log"
)

func InitDatabase() {
	err := db.InitDatabase()
	if err != nil {
		log.Println("Init database failed!")
		panic(err)
	}
	log.Println("Init database success.")
}