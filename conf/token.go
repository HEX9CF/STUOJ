package conf

type TokenConf struct {
	Expire int    `json:"expire"`
	Secret string `json:"secret"`
}
