package conf

type DatabaseConf struct {
	Host    string `yaml:"host" json:"host"`
	Port    string `yaml:"port" json:"port"`
	Name    string `yaml:"name" json:"name"`
	User    string `yaml:"user" json:"user"`
	Pwd     string `yaml:"password" json:"password"`
	MaxConn int    `yaml:"max_open_conns" json:"max_open_conns"`
	MaxIdle int    `yaml:"max_idle_conns" json:"max_idle_conns"`
}
