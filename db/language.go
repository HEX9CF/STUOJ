package db

import (
	"STUOJ/model"
	"log"
)

// 插入语言
func InsertLanguage(l model.Language) (uint64, error) {
	sql := "INSERT INTO tbl_language (id, name) VALUES (?, ?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(l.Id, l.Name)
	log.Println(sql, l.Id, l.Name)
	if err != nil {
		return 0, err
	}

	// 获取插入ID
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(id), nil
}

// 删除所有语言
func DeleteAllLanguages() error {
	sql := "DELETE FROM tbl_language"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	log.Println(sql)
	if err != nil {
		return err
	}

	return nil
}

// 查询所有语言
func SelectAllLanguages() ([]model.Language, error) {
	sql := "SELECT id, name FROM tbl_language"
	rows, err := db.Query(sql)
	log.Println(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// 遍历查询结果
	languages := make([]model.Language, 0)
	for rows.Next() {
		var language model.Language

		err := rows.Scan(&language.Id, &language.Name)
		if err != nil {
			return nil, err
		}

		//log.Println(language)
		languages = append(languages, language)
	}
	return languages, nil
}
