package yuki

import (
	"STUOJ/model"
	"encoding/json"
	"errors"
)

func GetAlbumList() ([]model.YukiAlbum, error) {
	bodystr, err := httpInteraction("/album", "GET", nil)
	if err != nil {
		return nil, err
	}

	var responses model.YukiResponses
	err = json.Unmarshal([]byte(bodystr), &responses)
	if err != nil {
		return nil, err
	}
	if responses.Code == 0 {
		return nil, errors.New(responses.Message)
	}
	if albumList, ok := responses.Data.([]model.YukiAlbum); ok {
		return albumList, nil
	}
	return nil, errors.New("albumList type assertion failed")
}

func GetAlbum(albumId uint64) (model.YukiAlbum, error) {
	bodystr, err := httpInteraction("/album/"+string(albumId), "GET", nil)
	if err != nil {
		return model.YukiAlbum{}, err
	}
	var responses model.YukiResponses
	err = json.Unmarshal([]byte(bodystr), &responses)
	if err != nil {
		return model.YukiAlbum{}, err
	}
	if responses.Code == 0 {
		return model.YukiAlbum{}, errors.New(responses.Message)
	}
	if album, ok := responses.Data.(model.YukiAlbum); ok {
		return album, nil
	}
	return model.YukiAlbum{}, errors.New("album type assertion failed")
}
