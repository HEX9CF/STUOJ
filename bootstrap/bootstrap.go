package bootstrap

func Init() {
	chFin := make(chan string)

	InitConfig()
	InitDatabase()

	// 异步初始化
	go InitJudge(chFin)
	go InitYuki(chFin)

	InitHandlers()
}
