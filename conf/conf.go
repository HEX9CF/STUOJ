package conf

import(
)

type Config struct{
	DateBase DatabaseConfig
}

// DatabaseConfig 数据库配置初始化
func InitConfig(){
	InitEnv()
	Conf=DefaultConfig()
}

// DefaultConfig 初始化Config并返回一个默认的Config指针
func DefaultConfig() *Config {
	return &Config{
		// Database
		DateBase: DatabaseConfigFromEnv(),
	}
}