package conf

import "os"

type ServerConf struct {
	Port string
}

func ServerConfigFromEnv() ServerConf {
	return ServerConf{
		Port: os.Getenv("SERVER_PORT"),
	}
}
