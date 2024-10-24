package conf

type DatabaseConf struct {
	Host    string `yaml:"host"`
	Port    string `yaml:"port"`
	Name    string `yaml:"name"`
	User    string `yaml:"user"`
	Pwd     string `yaml:"password"`
	MaxConn int    `yaml:"max_open_conns"`
	MaxIdle int    `yaml:"max_idle_conns"`
}
