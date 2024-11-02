package user_query

import (
	"STUOJ/db"
	"STUOJ/model"
	"html"
	"log"
	"strings"
	"time"
)

// 根据ID更新用户
func UpdateUserById(u model.User) error {
	// 预处理
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	err := u.HashPassword()
	if err != nil {
		return err
	}

	sql := "UPDATE tbl_user SET username = ?, email = ?, password = ?, avatar = ?, signature = ?, update_time = ? WHERE id = ?"
	stmt, err := db.Mysql.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	// 获取当前时间
	updateTime := time.Now().Format("2006-01-02 15:04:05")
	_, err = stmt.Exec(u.Username, u.Email, u.Password, u.Avatar, u.Signature, updateTime, u.Id)
	log.Println(sql, u.Username, u.Email, u.Password, u.Avatar, u.Signature, u.Password, updateTime, u.Id)
	if err != nil {
		return err
	}

	return nil
}

// 根据ID更新用户（除了密码）
func UpdateUserByIdExceptPassword(u model.User) error {
	// 预处理
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	sql := "UPDATE tbl_user SET username = ?, email = ?, signature = ?, update_time = ? WHERE id = ?"
	stmt, err := db.Mysql.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	// 获取当前时间
	updateTime := time.Now().Format("2006-01-02 15:04:05")
	_, err = stmt.Exec(u.Username, u.Email, u.Signature, updateTime, u.Id)
	log.Println(sql, u.Username, u.Email, u.Signature, updateTime, u.Id)
	if err != nil {
		return err
	}

	return nil
}

// 根据ID更新用户密码
func UpdateUserPasswordById(u model.User) error {
	// 预处理
	err := u.HashPassword()
	if err != nil {
		return err
	}

	sql := "UPDATE tbl_user SET password = ?, update_time = ? WHERE id = ?"
	stmt, err := db.Mysql.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	// 获取当前时间
	createTime := time.Now().Format("2006-01-02 15:04:05")
	updateTime := createTime
	_, err = stmt.Exec(u.Password, updateTime, u.Id)
	log.Println(sql, u.Password, updateTime, u.Id)
	if err != nil {
		return err
	}

	return nil
}

// 根据ID更新用户角色
func UpdateUserRoleById(u model.User) error {
	sql := "UPDATE tbl_user SET role = ? WHERE id = ?"
	stmt, err := db.Mysql.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(u.Role, u.Id)
	log.Println(sql, u.Role, u.Id)
	if err != nil {
		return err
	}

	return nil
}
