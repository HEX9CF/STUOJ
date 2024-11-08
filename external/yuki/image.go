package yuki

import (
	"STUOJ/internal/model"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"

	"github.com/mitchellh/mapstructure"
)

func UploadImage(path string, role uint8) (model.YukiImage, error) {
	url := preUrl + "/image"
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return model.YukiImage{}, err
	}
	req.Header.Set("Authorization", "Bearer "+config.Token)
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	err = writer.WriteField("album_name", model.GetAlbumName(role))
	if err != nil {
		return model.YukiImage{}, err
	}
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
	log.Println("resp.StatusCode: ", resp.StatusCode)
	bodys, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.YukiImage{}, err
	}
	bodystr := string(bodys)
	var responses model.YukiResponses
	err = json.Unmarshal([]byte(bodystr), &responses)
	if err != nil {
		return model.YukiImage{}, err
	}
	var image model.YukiImage
	err = mapstructure.Decode(responses.Data, &image)
	if err != nil {
		return model.YukiImage{}, err
	}
	return image, nil
}

func GetImageList(page uint64, role uint8) (model.YukiImageList, error) {
	bodystr, err := httpInteraction("/album/image/"+model.GetAlbumName(role)+"/?page="+strconv.FormatUint(page, 10), "GET", nil)
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
	var imageList model.YukiImageList
	err = mapstructure.Decode(responses.Data, &imageList)
	if err != nil {
		return model.YukiImageList{}, err
	}
	return imageList, nil
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
	var image model.YukiImage
	err = mapstructure.Decode(responses.Data, &image)
	if err != nil {
		return model.YukiImage{}, err
	}
	return image, nil
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
	var image model.YukiImage
	err = mapstructure.Decode(responses.Data, &image)
	if err != nil {
		return model.YukiImage{}, err
	}
	return image, nil
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
