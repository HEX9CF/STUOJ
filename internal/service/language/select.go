package language

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
)

// 查询所有语言
func SelectAll() ([]entity.Language, error) {
	var languages []entity.Language

	languages, err := dao.SelectAllLanguages()
	if err != nil {
		return nil, err
	}

	return languages, nil
}

// 根据名字模糊查询语言
func SelectLikeName(name string) (entity.Language, error) {
	var l entity.Language

	l, err := dao.SelectLanguageLikeName(name)
	if err != nil {
		return entity.Language{}, err
	}

	return l, nil
}
