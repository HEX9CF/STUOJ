package judge

import (
	"net/http"
)

func About() (*http.Response,error){
	url:=preUrl+"/about"
	req,_:=http.NewRequest("GET",url,nil)
	req.Header.Set("X-Auth-Token",config.Token)
	res,_:=http.DefaultClient.Do(req)
	defer res.Body.Close()
	return res,nil
}