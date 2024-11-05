package db

import (
	"STUOJ/model"
	"log"
)

// 插入标签
func InsertTag(t model.Tag) (uint64, error) {
	sql := "INSERT INTO tbl_tag (name) VALUES (?)"
	stmt, err := SqlDb.Prepare(sql)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(t.Name)
	log.Println(sql, t.Name)
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

// 根据ID查询标签
func SelectTagById(id uint64) (model.Tag, error) {
	var t model.Tag

	t.Id = id

	sql := "SELECT name FROM tbl_tag WHERE id = ?"
	err := SqlDb.QueryRow(sql, id).Scan(&t.Name)
	log.Println(sql, id)
	if err != nil {
		return model.Tag{}, err
	}

	return t, nil
}

// 查询所有标签
func SelectAllTags() ([]model.Tag, error) {
	sql := "SELECT id, name FROM tbl_tag"
	rows, err := SqlDb.Query(sql)
	log.Println(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// 遍历查询结果
	tags := make([]model.Tag, 0)
	for rows.Next() {
		var tag model.Tag

		err := rows.Scan(&tag.Id, &tag.Name)
		if err != nil {
			return nil, err
		}

		//log.Println(tag)
		tags = append(tags, tag)
	}
	return tags, nil
}

// 根据ID更新标签
func UpdateTagById(t model.Tag) error {
	sql := "UPDATE tbl_tag SET name = ? WHERE id = ?"
	stmt, err := SqlDb.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(t.Name, t.Id)
	log.Println(sql, t.Name, t.Id)
	if err != nil {
		return err
	}

	return nil
}

// 根据ID删除标签
func DeleteTagById(id uint64) error {
	sql := "DELETE FROM tbl_tag WHERE id = ?"
	stmt, err := SqlDb.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	log.Println(sql, id)
	if err != nil {
		return err
	}

	return nil
}
