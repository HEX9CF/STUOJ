package conf

type LskyproConf struct {
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
	ProblemToken string `yaml:"problem_token"`
	AvatarToken  string `yaml:"avatar_token"`
}
