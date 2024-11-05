package user_query

import (
	"STUOJ/db"
	"STUOJ/model"
	"log"
	"time"
)

// 根据ID查询用户
func SelectUserById(id uint64) (model.User, error) {
	var user model.User
	var createTimeStr, updateTimeStr string
	sql := "SELECT id, username, role, email, avatar, signature, create_time, update_time FROM tbl_user WHERE id = ? LIMIT 1"
	err := db.SqlDb.QueryRow(sql, id).Scan(&user.Id, &user.Username, &user.Role, &user.Email, &user.Avatar, &user.Signature, &createTimeStr, &updateTimeStr)
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
	sql := "SELECT id, username, role, email, avatar, signature, create_time, update_time FROM tbl_user"
	rows, err := db.SqlDb.Query(sql)
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

		err := rows.Scan(&user.Id, &user.Username, &user.Role, &user.Email, &user.Avatar, &user.Signature, &createTimeStr, &updateTimeStr)
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

// 根据角色ID查询用户
func SelectUsersByRole(r model.UserRole) ([]model.User, error) {
	sql := "SELECT id, username, role, email, avatar, signature, create_time, update_time FROM tbl_user WHERE role = ?"
	rows, err := db.SqlDb.Query(sql, r)
	log.Println(sql, r)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// 遍历查询结果
	users := make([]model.User, 0)
	for rows.Next() {
		var user model.User
		var createTimeStr, updateTimeStr string

		err := rows.Scan(&user.Id, &user.Username, &user.Role, &user.Email, &user.Avatar, &user.Signature, &createTimeStr, &updateTimeStr)
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
