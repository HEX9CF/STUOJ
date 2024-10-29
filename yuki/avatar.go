package yuki

import "STUOJ/model"

func UploadAvatar(path string) (model.YukiImage, error) {
	return UploadImage(path, model.YukiAvatarAlbum)

}
