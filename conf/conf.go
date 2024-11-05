package conf

import (
	"STUOJ/utils"
)

type Config struct {
	Datebase  DatabaseConf `yaml:"database" json:"database"`
	Judge     JudgeConf    `yaml:"judge" json:"judge"`
	YukiImage YukiConf     `yaml:"yuki-image" json:"yuki_image"`
	Server    ServerConf   `yaml:"server" json:"server"`
	Token     TokenConf    `yaml:"token" json:"token"`
}

// Config 初始化
func InitConfig() error {
	err := utils.ReadYaml(&Conf, "config.yaml")
	if err != nil {
		return err
	}
	utils.Expire = Conf.Token.Expire
	utils.Secret = Conf.Token.Secret
	utils.Refresh = Conf.Token.Refresh
	return nil
}
