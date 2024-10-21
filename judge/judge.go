package judge

import (
	"STUOJ/conf"
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	config conf.JudgeConf
	preUrl string
)

func InitJudge() error {
	config = conf.Conf.Judge
	preUrl = config.Host + ":" + config.Port
	response, err := About()
	if err != nil || response.StatusCode != http.StatusOK {
		log.Println("Judge server is not available!")
		return err
	}

	log.Println("Judge server is available.")
	return nil
}

func httpInteraction(route string, httpMethod string, reader *bytes.Reader) (string, error) {
	url := preUrl + route
	var req *http.Request
	var err error
	if reader == nil {
		req, err = http.NewRequest(httpMethod, url, nil)
	} else {
		req, err = http.NewRequest(httpMethod, url, reader)
	}
	if err != nil {
		return "", err
	}
	req.Header.Set("X-Auth-Token", config.Token)
	req.Header.Set("X-Auth-User", config.Token)
	req.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	bodystr := string(body)
	return bodystr, nil
}
