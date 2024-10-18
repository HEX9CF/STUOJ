package judge

import(
	"STUOJ/conf"
	"net/http"
	"fmt"
)
var(
	config conf.JudgeConf
	preUrl string
)
func InitJudge(){
	config=conf.Conf.Judge
	preUrl=config.Host+":"+config.Port
	response,err:=About()
	if err!=nil || response.StatusCode!=http.StatusOK{
		fmt.Println("Judge server is not available!"+err.Error())
	}else{
		fmt.Println("Judge server is available.")
	}
}