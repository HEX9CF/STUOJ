package conf

import (
	"os"
	"strconv"
)

type TokenConf struct {
	Expire int
	Secret string
}

func TokenConfigFromEnv() TokenConf {
	expire, _ := strconv.Atoi(os.Getenv("TOKEN_EXPIRE"))

	return TokenConf{
		Expire: expire,
		Secret: os.Getenv("API_SECRET"),
	}
}
