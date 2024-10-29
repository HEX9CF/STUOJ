package user_query

import (
	"STUOJ/database"
	"log"
	"time"
)

// 更新用户头像
func UpdateUserAvatar(id uint64, avatarUrl string) error {
	sql := "UPDATE tbl_user SET avatar=? ,update_time=? WHERE id=?"
	stmt, err := database.Db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	createTime := time.Now().Format("2006-01-02 15:04:05")
	updateTime := createTime
	_, err = stmt.Exec(avatarUrl, updateTime, id)
	log.Println(sql, avatarUrl, updateTime, id)
	if err != nil {
		return err
	}
	return nil
}

// 查询用户头像
func QueryUserAvatar(id uint64) (string, error) {
	var avatar string
	sql := "SELECT avatar FROM tbl_user WHERE id=?"
	err := database.Db.QueryRow(sql, id).Scan(&avatar)
	log.Println(sql, id)
	if err != nil {
		return "", err
	}
	return avatar, nil
}
