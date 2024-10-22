package lskypro

import (
	"STUOJ/model"
	"encoding/json"
)

func GetProfile() (model.LskyproProfile, error) {
	bodystr, err := httpInteraction("/profile", "GET", nil)
	if err != nil {
		return model.LskyproProfile{}, err
	}
	var profile model.LskyproProfile
	err = json.Unmarshal([]byte(bodystr), &profile)
	if err != nil {
		return model.LskyproProfile{}, err
	}
	return profile, nil
}
