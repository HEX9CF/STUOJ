package conf

import (
	_ "github.com/joho/godotenv"
	"os"
	"strconv"
)

type DatabaseConf struct {
	Host    string
	Port    string
	Name    string
	User    string
	Pwd     string
	MaxConn int
	MaxIdle int
}

// DatabaseConfigFromEnv 从环境变量中获取数据库配置信息，并返回DatabaseConfig结构体
func DatabaseConfigFromEnv() DatabaseConf {

	MaxConn, _ := strconv.Atoi(os.Getenv("DB_MAXOPENCONNS"))
	MaxIdle, _ := strconv.Atoi(os.Getenv("DB_MAXIDLECONNS"))

	return DatabaseConf{
		Host:    os.Getenv("DB_HOST"),
		Port:    os.Getenv("DB_PORT"),
		Name:    os.Getenv("DB_NAME"),
		User:    os.Getenv("DB_USER"),
		Pwd:     os.Getenv("DB_PASSWORD"),
		MaxConn: MaxConn,
		MaxIdle: MaxIdle,
	}
}
