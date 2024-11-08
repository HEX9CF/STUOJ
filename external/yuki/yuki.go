package yuki

import (
	"STUOJ/internal/conf"
	"bytes"
	"io"
	"log"
	"net/http"
)

func InitYukiImage() error {
	config = conf.Conf.YukiImage
	preUrl = config.Host + ":" + config.Port + "/api/v1"
	_, err := GetAlbumList()
	if err != nil {
		return err
	}
	log.Println("Successfully connected to yuki-image.")
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

	req.Header.Set("Authorization", "Bearer "+config.Token)
	req.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	bodystr := string(body)
	return bodystr, nil
}
