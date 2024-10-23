package lskypro

import (
	"STUOJ/model"
	"encoding/json"
	"errors"
)

func GetProfile(role uint8) (model.LskyproProfile, error) {
	bodystr, err := httpInteraction("/profile", "GET", nil, role)
	if err != nil {
		return model.LskyproProfile{}, err
	}
	var responses model.LskyproProfileResponses
	err = json.Unmarshal([]byte(bodystr), &responses)
	if err != nil {
		return model.LskyproProfile{}, err
	}
	if responses.Status == false {
		return model.LskyproProfile{}, errors.New(responses.Message)
	}
	return responses.Data, nil
}
