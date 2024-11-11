package bootstrap

func Init() {
	InitConfig()
	InitDatabase()

	// 异步初始化
	go InitJudge0()
	go InitYuki()

	InitServer()
}
