package bootstrap

func Init() {
	InitConfig()
	InitDatabase()

	// 异步初始化
	go InitJudge()
	go InitYuki()

	InitHandlers()
}
