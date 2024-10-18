package judge

import (
	"STUOJ/conf"
	"net/http"
	"bytes"
	"io/ioutil"
	"log"
	"STUOJ/model"
)

var (
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

func httpInteraction(route string,httpMethod string,reader *bytes.Reader)(string,error){
	url:=preUrl+route
	var req *http.Request
	var err error
	if reader==nil{
		req,err=http.NewRequest(httpMethod,url,nil)
	}else{
		req,err=http.NewRequest(httpMethod,url,reader)
	}
	if err != nil {
		return "",err
	}
	req.Header.Set("X-Auth-Token",config.Token)
	req.Header.Set("Content-Type","application/json")
	res,err:=http.DefaultClient.Do(req)
	if err != nil {
		return "",err
	}
	defer res.Body.Close()
	body,err:=ioutil.ReadAll(res.Body)
	if err != nil {
		return "",err
	}
	bodystr:=string(body)
	return bodystr,nil
}