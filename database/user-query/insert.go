package user_query

import (
	"STUOJ/database"
	"STUOJ/model"
	"html"
	"log"
	"strings"
	"time"
)

// 插入用户
func InsertUser(u model.User) (uint64, error) {
	// 预处理
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	err := u.HashPassword()
	if err != nil {
		return 0, err
	}

	sql := "INSERT INTO tbl_user (username, password, email, role, avatar, signature, create_time, update_time) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	stmt, err := database.Db.Prepare(sql)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	// 获取当前时间
	createTime := time.Now().Format("2006-01-02 15:04:05")
	updateTime := createTime
	result, err := stmt.Exec(u.Username, u.Password, u.Email, u.Role, u.Avatar, u.Signature, createTime, updateTime)
	log.Println(sql, u.Username, u.Password, u.Email, u.Role, u.Avatar, u.Signature, createTime, updateTime)
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

// 插入用户（注册）
func InsertUserForRegister(u model.User) (uint64, error) {
	// 预处理
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	err := u.HashPassword()
	if err != nil {
		return 0, err
	}

	// 默认值
	u.Avatar = "http://example.com/avatar.png"
	u.Signature = "这个大佬很懒，什么也没有留下"

	sql := "INSERT INTO tbl_user (username, password, email, avatar, signature, create_time, update_time) VALUES (?, ?, ?, ?, ?, ?, ?)"
	stmt, err := database.Db.Prepare(sql)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	// 获取当前时间
	createTime := time.Now().Format("2006-01-02 15:04:05")
	updateTime := createTime
	result, err := stmt.Exec(u.Username, u.Password, u.Email, u.Avatar, u.Signature, createTime, updateTime)
	log.Println(sql, u.Username, u.Password, u.Email, u.Avatar, u.Signature, createTime, updateTime)
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
