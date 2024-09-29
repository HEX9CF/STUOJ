package judge

import (
	"net/http"
)

func About() (*http.Response,error){
	return http.Get(config.Host+":"+config.Port)
}