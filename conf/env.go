package conf

import (
	"github.com/joho/godotenv"
	"log"
)

// 读取.env文件
func InitEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	log.Println("Loaded .env file")
	return nil
}
