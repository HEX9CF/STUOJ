package judge

import (
	conf2 "STUOJ/internal/conf"
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	config conf2.JudgeConf
	preUrl string
)

func InitJudge() error {
	config = conf2.Conf.Judge
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
	if route == "/submissions" && httpMethod == "POST" {
		log.Println("Wait for judge server to finish checking...")
		url = url + "?wait=true"
	}
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
	if res.StatusCode != http.StatusCreated && res.StatusCode != http.StatusOK {
		return "", errors.New(bodystr)
	}
	return bodystr, nil
}
