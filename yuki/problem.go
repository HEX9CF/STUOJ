package yuki

import "STUOJ/model"

func UpdateProblemImage(path string) (model.YukiImage, error) {
	return UploadImage(path, model.YukiProblemAlbum)
}
