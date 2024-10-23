package model

const (
	RoleProblem uint8 = 1
	RoleAvatar  uint8 = 2
)

type LskyproDeleteResponses struct {
	Status  bool                   `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

type LskyproProfileResponses struct {
	Status  bool           `json:"status"`
	Message string         `json:"message"`
	Data    LskyproProfile `json:"data"`
}

type LskyproProfile struct {
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
	Size       float64                     `json:"size"`
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

type LskyproImageListResponses struct {
	Status  bool             `json:"status"`
	Message string           `json:"message"`
	Data    LskyproImageList `json:"data"`
}

type LskyproImageList struct {
	CurrentPage  uint64                 `json:"current_page"`
	Data         []LskyproImageData     `json:"data"`
	FirstPageUrl string                 `json:"first_page_url"`
	From         uint64                 `json:"from"`
	LastPage     uint64                 `json:"last_page"`
	LastPageUrl  string                 `json:"last_page_url"`
	Links        []LskyproImageListLink `json:"links"`
	NextPageUrl  string                 `json:"next_page_url"`
	Path         string                 `json:"path"`
	PerPage      uint64                 `json:"per_page"`
	PrevPageUrl  uint64                 `json:"prev_page_url"`
	To           uint64                 `json:"to"`
	Total        uint64                 `json:"total"`
}

type LskyproImageListLink struct {
	Url    string `json:"url"`
	Label  string `json:"label"`
	Active bool   `json:"active"`
}
type LskyproImageData struct {
	Key        string               `json:"key,omitempty"`
	Name       string               `json:"name,omitempty"`
	OriginName string               `json:"origin_name,omitempty"`
	Size       float64              `json:"size,omitempty"`
	Mimetype   string               `json:"mimetype,omitempty"`
	Extension  string               `json:"extension,omitempty"`
	Md5        string               `json:"md5,omitempty"`
	Sha1       string               `json:"sha1,omitempty"`
	Width      uint64               `json:"width,omitempty"`
	Height     uint64               `json:"height,omitempty"`
	HumanDate  string               `json:"human_date,omitempty"`
	Date       string               `json:"date,omitempty"`
	Pathname   string               `json:"pathname,omitempty"`
	Links      LskyproImageDataLink `json:"links,omitempty"`
}

type LskyproImageDataLink struct {
	Url              string `json:"url"`
	Html             string `json:"html"`
	Bbcode           string `json:"bbcode"`
	Markdown         string `json:"markdown"`
	MarkdownWithLink string `json:"markdown_with_link"`
	ThumbnailUrl     string `json:"thumbnail_url"`
}
