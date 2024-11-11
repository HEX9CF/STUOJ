package conf

import (
	"STUOJ/utils"
)

type Config struct {
	Datebase  DatabaseConf `yaml:"database" json:"database"`
	Judge     JudgeConf    `yaml:"judge0" json:"judge0"`
	YukiImage YukiConf     `yaml:"yuki-image" json:"yuki_image"`
	Server    ServerConf   `yaml:"server" json:"server"`
	Token     TokenConf    `yaml:"token" json:"token"`
}

// Config 初始化
func InitConfig() error {
	v, err := utils.IsFileExists("config.yaml")
	if err != nil {
		return err
	}
	if !v {
		Conf.Default()
		utils.WriteYaml(&Conf, "config.yaml")
	}
	err = utils.ReadYaml(&Conf, "config.yaml")
	if err != nil {
		return err
	}
	utils.Expire = Conf.Token.Expire
	utils.Secret = Conf.Token.Secret
	utils.Refresh = Conf.Token.Refresh
	return nil
}

func (c *Config) Default() {
	c.Datebase.Default()
	c.Judge.Default()
	c.YukiImage.Default()
	c.Server.Default()
	c.Token.Default()
}
