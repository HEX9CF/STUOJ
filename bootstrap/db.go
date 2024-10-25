package bootstrap

import (
	"STUOJ/db"
	"log"
)

func InitDatabase(chFin chan string) {
	err := db.InitDatabase()
	if err != nil {
		log.Println("Init database failed!")
		panic(err)
	}
	chFin <- "database"
}
