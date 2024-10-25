package bootstrap

import "log"

func Init() {
	chFin := make(chan string)

	// 优先初始化配置
	InitConfig()

	// 异步初始化
	go InitDatabase(chFin)
	go InitJudge(chFin)
	go InitLskypro(chFin)

	// 等待其他初始化完成
	const bootstraps = 3
	count := 0
	for {
		select {
		case name := <-chFin:
			log.Println("Init", name, "success.")
			count++
		}
		if count == bootstraps {
			break
		}
	}

	// 最后初始化路由
	InitHandlers()
}
