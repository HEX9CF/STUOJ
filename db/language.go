package db

import (
	"STUOJ/model"
)

// 插入语言
func InsertLanguage(l model.Language) (uint64, error) {
	tx := Db.Create(&l)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return l.Id, nil
}

// 查询所有语言
func SelectAllLanguages() ([]model.Language, error) {
	var languages []model.Language

	tx := Db.Find(&languages)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return languages, nil
}

// 根据名字模糊查询语言
func SelectLanguageLikeName(name string) (model.Language, error) {
	var l model.Language

	tx := Db.Where("name like ?", "%"+name+"%").First(&l)
	if tx.Error != nil {
		return model.Language{}, tx.Error
	}

	return l, nil
}

// 删除所有语言
func DeleteAllLanguages() error {
	tx := Db.Where("1 = 1").Delete(&model.Language{})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
