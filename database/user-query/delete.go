package user_query

import (
	"STUOJ/database"
	"log"
)

// 根据ID删除用户
func DeleteUserById(id uint64) error {
	sql := "DELETE FROM tbl_user WHERE id = ?"
	stmt, err := database.Db.Prepare(sql)
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
