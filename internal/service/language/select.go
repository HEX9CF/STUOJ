package language

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"errors"
	"log"
)

// 查询所有语言
func SelectAll() ([]entity.Language, error) {
	var languages []entity.Language

	languages, err := dao.SelectAllLanguages()
	if err != nil {
		log.Println(err)
		return nil, errors.New("查询语言失败")
	}

	return languages, nil
}

// 根据名字模糊查询语言
func SelectLikeName(name string) (entity.Language, error) {
	var l entity.Language

	l, err := dao.SelectLanguageLikeName(name)
	if err != nil {
		log.Println(err)
		return entity.Language{}, errors.New("查询语言失败")
	}

	return l, nil
}
