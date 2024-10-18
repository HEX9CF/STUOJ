package judge

import(
	"encoding/json"
	"net/http"
	"log"
	"io/ioutil"
)

func GetLanguage()([]map[string]interface{}){
	url:=preUrl+"/languages"
	req,_:=http.NewRequest("GET",url,nil)
	req.Header.Set("X-Auth-Token",config.Token)
	res,_:=http.DefaultClient.Do(req)
	defer res.Body.Close()
	body,_:=ioutil.ReadAll(res.Body)
	bodystr:=string(body)
	var dataArr []map[string]interface{}
	err:= json.Unmarshal([]byte(bodystr),&dataArr)
	if err!=nil{
		log.Println(err)
	}
	return dataArr
}