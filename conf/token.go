package conf

type TokenConf struct {
	Expire  uint64 `yaml:"expire" json:"expire"`
	Refresh uint64 `yaml:"refresh" json:"refresh"`
	Secret  string `yaml:"secret" json:"secret"`
}
