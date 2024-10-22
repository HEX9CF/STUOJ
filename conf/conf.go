package conf

type Config struct {
	DateBase DatabaseConf
	Judge    JudgeConf
	Lskypro  LskyproConf
	Server   ServerConf
	Token    TokenConf
}

// Config 初始化
func InitConfig() error {
	err := InitEnv()
	if err != nil {
		return err
	}
	Conf = DefaultConfig()

	return nil
}

// DefaultConfig 初始化Config并返回一个默认的Config指针
func DefaultConfig() *Config {
	return &Config{
		// Database
		DateBase: DatabaseConfigFromEnv(),
		// Judge
		Judge: JudgeConfigFromEnv(),
		// Lskypro
		Lskypro: LskyproConfigFromEnv(),
		// Server
		Server: ServerConfigFromEnv(),
		// Token
		Token: TokenConfigFromEnv(),
	}
}
