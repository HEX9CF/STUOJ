package bootstrap

import (
	"STUOJ/conf"
	"STUOJ/handlers"
)

func Init() {
	conf.InitConfig()
	InitDatabase()
	InitJudge()
	InitHandlers()
}
