package conf

import (
	"github.com/joho/godotenv"
	"log"
)

// 读取.env文件
func InitEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Loaded .env file")
}
