package db

import (
	"STUOJ/model"
	"log"
)

// 插入语言
func InsertLanguage(l model.Language) error {
	sql := "INSERT INTO tbl_language (id, name) VALUES (?, ?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(l.Id, l.Name)
	log.Println(sql, l.Id, l.Name)
	if err != nil {
		return err
	}

	return nil
}

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
