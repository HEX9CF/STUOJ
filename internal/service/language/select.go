package language

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/model"
)

// 查询所有语言
func SelectAll() ([]model.Language, error) {
	var languages []model.Language

	languages, err := dao.SelectAllLanguages()
	if err != nil {
		return nil, err
	}

	return languages, nil
}

// 根据名字模糊查询语言
func SelectLikeName(name string) (model.Language, error) {
	var l model.Language

	l, err := dao.SelectLanguageLikeName(name)
	if err != nil {
		return model.Language{}, err
	}

	return l, nil
}
