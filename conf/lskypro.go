package conf

import (
	"os"

	_ "github.com/joho/godotenv"
)

type LskyproConf struct {
	Host  string
	Port  string
	Token string
}

func LskyproConfigFromEnv() LskyproConf {
	return LskyproConf{
		Host:  os.Getenv("LSKYPRO_HOST"),
		Port:  os.Getenv("LSKYPRO_PORT"),
		Token: os.Getenv("LSKYPRO_TOKEN"),
	}
}
