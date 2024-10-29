package conf

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	DateBase  DatabaseConf `yaml:"db"`
	Judge     JudgeConf    `yaml:"judge"`
	YukiImage YukiConf     `yaml:"yuki-image"`
	Server    ServerConf   `yaml:"server"`
	Token     TokenConf    `yaml:"token"`
}

// Config 初始化
func InitConfig() error {
	file, err := os.Open("config.yaml")
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	Conf = &Config{}
	err = decoder.Decode(Conf)
	if err != nil {
		return err
	}
	return nil
}
