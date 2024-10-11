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

	//row, err := db.Query("SELECT * FROM tbl_user")
	//if err != nil {
	//log.Println("Error querying the database:", err)
	//return nil
	//}
	////defer row.Close()
	//
	//for row.Next() {
	//var id int
	//var username string
	//var password string
	//var role int
	//var email string
	//var avatar string
	//var create_time string
	//var update_time string
	//err = row.Scan(&id, &username, &password, &role, &email, &avatar, &create_time, &update_time)
	//if err != nil {
	//log.Println("Error scanning the database:", err)
	//return nil
	//}
	//log.Println(id, username, password, role, email, avatar, create_time, update_time)
	//}
	//
	//return nil
}
