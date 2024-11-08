package fps

import (
	"STUOJ/internal/model"
	"encoding/xml"
	"io"
	"os"
)

func Read(path string) (model.FPS, error) {
	xmlFile, err := os.Open(path)
	if err != nil {
		return model.FPS{}, err
	}
	defer xmlFile.Close()

	// 读取文件内容
	xmlData, err := io.ReadAll(xmlFile)
	if err != nil {
		return model.FPS{}, err
	}

	// 解析XML
	var fps model.FPS
	err = xml.Unmarshal(xmlData, &fps)
	if err != nil {
		return model.FPS{}, err
	}
	return fps, nil
}
