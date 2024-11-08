package yuki

import (
	"STUOJ/internal/model"
)

func UpdateProblemImage(path string) (model.YukiImage, error) {
	return UploadImage(path, model.YukiProblemAlbum)
}
