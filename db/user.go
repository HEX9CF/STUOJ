package db

import (
	"STUOJ/model"
	"log"
)

func GetAllUsers() []model.User {
	rows, err := db.Query("SELECT * FROM tbl_user")
	log.Println("SELECT * FROM tbl_user")
	log.Println(rows)
	if err != nil {
		return nil
	}
	defer rows.Close()

	users := make([]model.User, 0)
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Role, &user.Email, &user.Avatar, &user.CreateTime, &user.UpdateTime)
		if err != nil {
			return nil
		}
		users = append(users, user)
	}
	return users
}
