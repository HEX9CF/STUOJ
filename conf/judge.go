package conf

import(
	_ "github.com/joho/godotenv"
	"os"
)

type JudgeConf struct {
	Host string
	Port string
	Token string
}

func JudgeConfigFromEnv() JudgeConf {
	return JudgeConf{
		Host: os.Getenv("JUDGE_HOST"),
		Port: os.Getenv("JUDGE_PORT"),
		Token: os.Getenv("JUDGE_TOKEN"),
	}
}