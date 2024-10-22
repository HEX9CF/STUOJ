package model

type LskyproProfile struct {
	Status  bool               `json:"status"`
	Message string             `json:"message"`
	Data    LskyproProfileData `json:"data"`
}

type LskyproProfileData struct {
	Name          string  `json:"name"`
	Avatar        string  `json:"avatar"`
	Email         string  `json:"email"`
	Capacity      float64 `json:"capacity"`
	Used_capacity float64 `json:"used_capacity"`
	Url           string  `json:"url"`
	Image_num     uint64  `json:"image_num"`
	Album_num     uint64  `json:"album_num"`
	Registered_ip string  `json:"registered_ip"`
}
