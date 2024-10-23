package lskypro

import (
	"STUOJ/model"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) (model.LskyproUploadData, error) {
	url := preUrl + "/upload"

	var fileJson model.UploadImageData
	if err := c.ShouldBind(&fileJson); err != nil {
		return model.LskyproUploadData{}, err
	}

	file, err := c.FormFile("file")
	if err != nil {
		return model.LskyproUploadData{}, err
	}

	// 保存文件到临时位置
	dst := fmt.Sprintf("tmp/%s", file.Filename)

	if err := c.SaveUploadedFile(file, dst); err != nil {
		return model.LskyproUploadData{}, err
	}

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return model.LskyproUploadData{}, err
	}
	if fileJson.Role == model.RoleProblem {
		req.Header.Set("Authorization", "Bearer "+config.ProblemToken)
	} else if fileJson.Role == model.RoleAvatar {
		req.Header.Set("Authorization", "Bearer "+config.AvatarToken)
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", file.Filename)
	if err != nil {
		return model.LskyproUploadData{}, err
	}

	src, err := os.Open(dst)
	if err != nil {
		return model.LskyproUploadData{}, err
	}
	defer src.Close()

	_, err = io.Copy(part, src)
	if err != nil {
		return model.LskyproUploadData{}, err
	}
	err = writer.Close()
	if err != nil {
		return model.LskyproUploadData{}, err
	}
	req.Body = io.NopCloser(body)
	req.ContentLength = int64(body.Len())
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// 如果发送请求失败，返回错误信息
		return model.LskyproUploadData{}, err
	}
	defer resp.Body.Close()

	bodys, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.LskyproUploadData{}, err
	}
	bodystr := string(bodys)
	var responses model.LskyproUploadResponses
	err = json.Unmarshal([]byte(bodystr), &responses)
	if err != nil {
		return model.LskyproUploadData{}, err
	}
	if responses.Status == false {
		return model.LskyproUploadData{}, errors.New(responses.Message)
	}
	return responses.Data, nil
}

func GetImageList(page uint64, role uint8) (model.LskyproImageList, error) {
	bodystr, err := httpInteraction("/images"+"/?page="+strconv.FormatUint(page, 10), "GET", nil, role)
	if err != nil {
		return model.LskyproImageList{}, err
	}
	var responses model.LskyproImageListResponses
	err = json.Unmarshal([]byte(bodystr), &responses)
	if err != nil {
		return model.LskyproImageList{}, err
	}
	if responses.Status == false {
		return model.LskyproImageList{}, errors.New(responses.Message)
	}
	return responses.Data, nil
}

func Delete(key string, role uint8) error {
	bodeystr, err := httpInteraction("/images"+"/"+key, "DELETE", nil, role)
	if err != nil {
		return err
	}
	var responses model.LskyproDeleteResponses
	err = json.Unmarshal([]byte(bodeystr), &responses)
	if err != nil {
		return err
	}
	if responses.Status == false {
		return errors.New(responses.Message)
	}
	return nil
}
