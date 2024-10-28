package db

import (
	"STUOJ/model"
	"html"
	"log"
	"strings"
	"time"
)

// 根据ID查询用户
func SelectUserById(id uint64) (model.User, error) {
	var user model.User
	var createTimeStr, updateTimeStr string
	sql := "SELECT id, username, role, email, avatar, signature, create_time, update_time FROM tbl_user WHERE id = ? LIMIT 1"
	err := db.QueryRow(sql, id).Scan(&user.Id, &user.Username, &user.Role, &user.Email, &user.Avatar, &user.Signature, &createTimeStr, &updateTimeStr)
	log.Println(sql, id)
	if err != nil {
		return model.User{}, err
	}

	// 时间格式转换
	timeLayout := "2006-01-02 15:04:05"
	user.CreateTime, err = time.Parse(timeLayout, createTimeStr)
	if err != nil {
		return model.User{}, err
	}
	user.UpdateTime, err = time.Parse(timeLayout, updateTimeStr)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

// 查询所有用户
func SelectAllUsers() ([]model.User, error) {
	sql := "SELECT id, username, role, email, avatar, create_time, update_time FROM tbl_user"
	rows, err := db.Query(sql)
	log.Println(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// 遍历查询结果
	users := make([]model.User, 0)
	for rows.Next() {
		var user model.User
		var createTimeStr, updateTimeStr string

		err := rows.Scan(&user.Id, &user.Username, &user.Role, &user.Email, &user.Avatar, &createTimeStr, &updateTimeStr)
		if err != nil {
			return nil, err
		}

		// 时间格式转换
		timeLayout := "2006-01-02 15:04:05"
		user.CreateTime, err = time.Parse(timeLayout, createTimeStr)
		if err != nil {
			return nil, err
		}
		user.UpdateTime, err = time.Parse(timeLayout, updateTimeStr)
		if err != nil {
			return nil, err
		}

		//log.Println(user)
		users = append(users, user)
	}
	return users, nil
}

// 插入用户
func InsertUser(u model.User) (uint64, error) {
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
	stmt, err := db.Prepare(sql)
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

// 根据ID更新用户（除了密码）
func UpdateUserById(u model.User) error {
	// 预处理
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	sql := "UPDATE tbl_user SET username = ?, email = ?, signature = ?, update_time = ? WHERE id = ?"
	stmt, err := db.Prepare(sql)
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
	stmt, err := db.Prepare(sql)
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

// 根据ID删除用户
func DeleteUserById(id uint64) error {
	sql := "DELETE FROM tbl_user WHERE id = ?"
	stmt, err := db.Prepare(sql)
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

// 根据邮箱验证密码
func VerifyUserByEmail(u model.User) (uint64, error) {
	//log.Println("用户登录：", u.Email, u.Password)

	// 查询用户
	var id uint64
	var hashedPassword string
	sql := "SELECT id, password FROM tbl_user WHERE email = ? LIMIT 1"
	err := db.QueryRow(sql, &u.Email).Scan(&id, &hashedPassword)
	log.Println(sql, u.Email)
	if err != nil {
		return 0, err
	}

	// 验证密码
	//log.Println("验证密码：", u.Password, hashedPassword)
	err = u.VerifyByHashedPassword(hashedPassword)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// 根据ID验证密码
func VerifyUserById(u model.User) (uint64, error) {
	//log.Println("用户登录：", u.Email, u.Password)

	// 查询用户
	var id uint64
	var hashedPassword string
	sql := "SELECT id, password FROM tbl_user WHERE id = ? LIMIT 1"
	err := db.QueryRow(sql, &u.Id).Scan(&id, &hashedPassword)
	log.Println(sql, u.Id)
	if err != nil {
		return 0, err
	}

	// 验证密码
	log.Println("验证密码：", u.Password, hashedPassword)
	err = u.VerifyByHashedPassword(hashedPassword)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// 更新用户头像
func UpdateUserAvatar(id uint64, avatarUrl string) error {
	sql := "UPDATE tbl_user SET avatar=? ,update_time=? WHERE id=?"
	stmt, err := db.Prepare(sql)
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
	err := db.QueryRow(sql, id).Scan(&avatar)
	log.Println(sql, id)
	if err != nil {
		return "", err
	}
	return avatar, nil
}
