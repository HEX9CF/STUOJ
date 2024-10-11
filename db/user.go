package db

import (
	"STUOJ/model"
	"log"
	"time"
)

func GetAllUsers() []model.User {
	// 查询所有用户
	sql := "SELECT * FROM tbl_user"
	rows, err := db.Query(sql)
	log.Println(sql)
	if err != nil {
		return nil
	}
	defer rows.Close()

	// 遍历查询结果
	users := make([]model.User, 0)
	for rows.Next() {
		var user model.User
		var createTimeStr, updateTimeStr string

		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Role, &user.Email, &user.Avatar, &createTimeStr, &updateTimeStr)
		if err != nil {
			log.Println(err)
			return nil
		}

		// 时间格式转换
		timeLayout := "2006-01-02 15:04:05"
		user.CreateTime, err = time.Parse(timeLayout, createTimeStr)
		if err != nil {
			log.Println(err)
			return nil
		}
		user.UpdateTime, err = time.Parse(timeLayout, updateTimeStr)
		if err != nil {
			log.Println(err)
			return nil
		}

		//log.Println(user)
		users = append(users, user)
	}
	return users
}
