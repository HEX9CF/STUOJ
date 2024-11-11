package yuki

import "STUOJ/internal/model"

func UploadAvatar(path string) (model.YukiImage, error) {
	return UploadImage(path, model.YukiAvatarAlbum)
}

func DeleteOldAvatar(url string) error {
	image, err := GetImageFromUrl(url)
	if err != nil {
		return err
	}
	return DeleteImage(image.Key)
}
