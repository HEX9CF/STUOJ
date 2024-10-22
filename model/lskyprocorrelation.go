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

type LskyproUploadResponses struct {
	Status  bool              `json:"status"`
	Message string            `json:"message"`
	Data    LskyproUploadData `json:"data"`
}

type LskyproUploadData struct {
	Key        string                      `json:"key"`
	Name       string                      `json:"name"`
	Pathname   string                      `json:"pathname"`
	OriginName string                      `json:"origin_name"`
	Size       float32                     `json:"size"`
	Mimetype   string                      `json:"mimetype"`
	Extension  string                      `json:"extension"`
	Md5        string                      `json:"md5"`
	Sha1       string                      `json:"sha1"`
	Links      LskyproUploadResponsesLinks `json:"links"`
}

type LskyproUploadResponsesLinks struct {
	Url              string `json:"url"`
	Html             string `json:"html"`
	Bbcode           string `json:"bbcode"`
	Markdown         string `json:"markdown"`
	MarkdownWithLink string `json:"markdown_with_link"`
	ThumbnailUrl     string `json:"thumbnail_url"`
}
