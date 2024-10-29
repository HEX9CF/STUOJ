package user_query

import (
	"STUOJ/database"
	"STUOJ/model"
	"log"
)

// 根据邮箱验证密码
func VerifyUserByEmail(u model.User) (uint64, error) {
	//log.Println("用户登录：", u.Email, u.Password)

	// 查询用户
	var id uint64
	var hashedPassword string
	sql := "SELECT id, password FROM tbl_user WHERE email = ? LIMIT 1"
	err := database.Db.QueryRow(sql, &u.Email).Scan(&id, &hashedPassword)
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
	err := database.Db.QueryRow(sql, &u.Id).Scan(&id, &hashedPassword)
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
