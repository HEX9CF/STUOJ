package conf

type LimiterConf struct {
	Second int `yaml:"second"`
	Size   int `yaml:"size"`
}
