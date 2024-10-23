package lskypro

import (
	"STUOJ/conf"
	"STUOJ/model"
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

func InitLskypro() error {
	config = conf.Conf.Lskypro
	preUrl = config.Host + ":" + config.Port + "/api/v1"
	_, err := GetProfile(1)
	if err != nil {
		return err
	}
	log.Println("Successfully connected to LskyPro.")
	return nil
}

func httpInteraction(route string, httpMethod string, reader *bytes.Reader, role uint8) (string, error) {
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
	if role == model.RoleProblem {
		req.Header.Set("Authorization", "Bearer "+config.ProblemToken)
	} else if role == model.RoleAvatar {
		req.Header.Set("Authorization", "Bearer "+config.AvatarToken)
	}
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
