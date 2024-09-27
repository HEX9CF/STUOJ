package conf

import(
)

type Config struct{
	DateBase DatabaseConfig
}

func InitConfig(){
	InitEnv()
	Conf=DefaultConfig()
}

func DefaultConfig() *Config {
	return &Config{
		// Database
		DateBase: DatabaseConfigFromEnv(),
	}
}