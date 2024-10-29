package yuki

import (
	"STUOJ/model"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
)

func UploadImage(path string, roal uint8) (model.YukiImage, error) {
	url := preUrl + "/image"
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return model.YukiImage{}, err
	}
	req.Header.Set("Authorization", "Bearer "+config.Token)
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fileinfo, err := os.Stat(path)
	part, err := writer.CreateFormFile("file", fileinfo.Name())
	if err != nil {
		return model.YukiImage{}, err
	}

	src, err := os.Open(path)
	if err != nil {
		return model.YukiImage{}, err
	}
	defer src.Close()

	_, err = io.Copy(part, src)
	if err != nil {
		return model.YukiImage{}, err
	}
	err = writer.Close()
	if err != nil {
		return model.YukiImage{}, err
	}
	req.Body = io.NopCloser(body)
	req.ContentLength = int64(body.Len())
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// 如果发送请求失败，返回错误信息
		return model.YukiImage{}, err
	}
	defer resp.Body.Close()

	bodys, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.YukiImage{}, err
	}
	bodystr := string(bodys)
	var responses model.YukiResponses
	err = json.Unmarshal([]byte(bodystr), &responses)
	if image, ok := responses.Data.(model.YukiImage); ok {
		return image, nil
	}
	return model.YukiImage{}, errors.New("Upload failed")
}

func GetImageList(page uint64, role uint8) (model.YukiImageList, error) {
	bodystr, err := httpInteraction("/album/image"+strconv.FormatUint(uint64(role), 10)+"/?page="+strconv.FormatUint(page, 10), "GET", nil)
	if err != nil {
		return model.YukiImageList{}, err
	}
	var responses model.YukiResponses
	err = json.Unmarshal([]byte(bodystr), &responses)
	if err != nil {
		return model.YukiImageList{}, err
	}
	if responses.Code == 0 {
		return model.YukiImageList{}, errors.New(responses.Message)
	}
	if imageList, ok := responses.Data.(model.YukiImageList); ok {
		return imageList, nil
	}
	return model.YukiImageList{}, errors.New("Get failed")
}

func GetImage(key string) (model.YukiImage, error) {
	bodystr, err := httpInteraction("/image/"+key, "GET", nil)
	if err != nil {
		return model.YukiImage{}, err
	}
	var responses model.YukiResponses
	err = json.Unmarshal([]byte(bodystr), &responses)
	if err != nil {
		return model.YukiImage{}, err
	}
	if responses.Code == 0 {
		return model.YukiImage{}, errors.New(responses.Message)
	}
	if image, ok := responses.Data.(model.YukiImage); ok {
		return image, nil
	}
	return model.YukiImage{}, errors.New("Get failed")
}

func GetImageFromUrl(url string) (model.YukiImage, error) {
	bodystr, err := httpInteraction("/image/?url="+url, "GET", nil)
	if err != nil {
		return model.YukiImage{}, err
	}
	var responses model.YukiResponses
	err = json.Unmarshal([]byte(bodystr), &responses)
	if err != nil {
		return model.YukiImage{}, err
	}
	if responses.Code == 0 {
		return model.YukiImage{}, errors.New(responses.Message)
	}
	if image, ok := responses.Data.(model.YukiImage); ok {
		return image, nil
	}
	return model.YukiImage{}, errors.New("Get failed")
}

func DeleteImage(key string) error {
	bodeystr, err := httpInteraction("/image"+"/"+key, "DELETE", nil)
	if err != nil {
		return err
	}
	var responses model.YukiResponses
	err = json.Unmarshal([]byte(bodeystr), &responses)
	if err != nil {
		return err
	}
	if responses.Code == 0 {
		return errors.New(responses.Message)
	}
	return nil
}
