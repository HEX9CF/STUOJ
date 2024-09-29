package bootstrap

import (
	"STUOJ/conf"
)
func Init(){
	conf.InitConfig()
	InitDatabase()
	InitJudge()
	InitHandlers()
}