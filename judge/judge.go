package judge

import (
	"STUOJ/conf"
	"log"
	"net/http"
)

var (
	config conf.JudgeConf
	preUrl string
)

func InitJudge() {
	config = conf.Conf.Judge
	preUrl = config.Host + ":" + config.Port
	response, err := About()
	if err != nil || response.StatusCode != http.StatusOK {
		log.Println("Judge server is not available!" + err.Error())
	} else {
		log.Println("Judge server is available.")
	}
}
