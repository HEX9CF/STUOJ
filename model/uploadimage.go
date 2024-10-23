package model

import "mime/multipart"

type UploadImageData struct {
	File *multipart.FileHeader `form:"file"`
	Role uint8                 `form:"role"`
}
