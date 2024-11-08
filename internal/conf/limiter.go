package conf

type LimiterConf struct {
	Second int `yaml:"Second"`
	Size   int `yaml:"size"`
}
