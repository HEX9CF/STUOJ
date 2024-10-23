package conf

import (
	"os"

	_ "github.com/joho/godotenv"
)

type LskyproConf struct {
	Host         string
	Port         string
	ProblemToken string
	AvatarToken  string
}

func LskyproConfigFromEnv() LskyproConf {
	return LskyproConf{
		Host:         os.Getenv("LSKYPRO_HOST"),
		Port:         os.Getenv("LSKYPRO_PORT"),
		ProblemToken: os.Getenv("LSKYPRO_PROBLEM_TOKEN"),
		AvatarToken:  os.Getenv("LSKYPRO_AVATAR_TOKEN"),
	}
}
