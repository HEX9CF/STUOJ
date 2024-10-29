package conf

type TokenConf struct {
	Expire  uint64 `json:"expire"`
	Refresh uint64 `json:"refresh"`
	Secret  string `json:"secret"`
}
