package conf

import (
	"fmt"
	"log"
	"github.com/joho/godotenv"
)

// 读取.env文件
func InitEnv(){
	err := godotenv.Load()
	if err != nil {
	  log.Fatal(err)
	}
	fmt.Println("Loaded .env file")
}

