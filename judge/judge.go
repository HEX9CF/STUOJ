package judge

import(
	"STUOJ/conf"
	"net/http"
	"log"
	"STUOJ/model"
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
		log.Println("Judge server is not available!"+err.Error())
	}else{
		log.Println("Judge server is available.")
	}
}

func test(){
	var submission model.JudgeSubmission
	submission=model.JudgeSubmission{
		SourceCode: "#include <stdio.h>\n\nint main(void) {\n  char name[10];\n  scanf(\"%s\", name);\n  printf(\"hello, %s\\n\", name);\n  return 0;\n}",
		LanguageId: 2,
		Stdin:"world",
		ExpectedOutput:"hello, world\n",
		CPUTimeLimit:  0.005,
		MemoryLimit:2048,
	}
	log.Println(Submit(submission))
}