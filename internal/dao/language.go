package dao

import (
	"STUOJ/internal/db"
	"STUOJ/internal/entity"
)

// 插入语言
func InsertLanguage(l entity.Language) (uint64, error) {
	tx := db.Db.Create(&l)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return l.Id, nil
}

// 查询所有语言
func SelectAllLanguages() ([]entity.Language, error) {
	var languages []entity.Language

	tx := db.Db.Find(&languages)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return languages, nil
}

// 根据名字模糊查询语言
func SelectLanguageLikeName(name string) (entity.Language, error) {
	var l entity.Language

	tx := db.Db.Where("name like ?", "%"+name+"%").First(&l)
	if tx.Error != nil {
		return entity.Language{}, tx.Error
	}

	return l, nil
}

// 删除所有语言
func DeleteAllLanguages() error {
	tx := db.Db.Where("1 = 1").Delete(&entity.Language{})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// 统计语言数量
func CountLanguages() (int64, error) {
	var count int64

	tx := db.Db.Model(&entity.Language{}).Count(&count)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return count, nil
}
