package conf

import (
	"fmt"
	"log"
	"github.com/joho/godotenv"
)

func InitEnv(){
	err := godotenv.Load()
	if err != nil {
	  log.Fatal(err)
	}
	fmt.Println("Loaded .env file")
}

