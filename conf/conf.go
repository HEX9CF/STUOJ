package conf

type Config struct {
	DateBase DatabaseConf
	Judge    JudgeConf
	Server   ServerConf
	Token    TokenConf
}

// Config 初始化
func InitConfig() {
	InitEnv()
	Conf = DefaultConfig()
}

// DefaultConfig 初始化Config并返回一个默认的Config指针
func DefaultConfig() *Config {
	return &Config{
		// Database
		DateBase: DatabaseConfigFromEnv(),
		// Judge
		Judge: JudgeConfigFromEnv(),
		// Server
		Server: ServerConfigFromEnv(),
		// Token
		Token: TokenConfigFromEnv(),
	}
}
