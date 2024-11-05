package db

import (
	"STUOJ/model"
	"html"
	"strings"
	"time"
)

// 根据ID查询用户
func SelectUserById(id uint64) (model.User, error) {
	var user model.User

	tx := Db.Where("id = ?", id).First(&user)
	if tx.Error != nil {
		return model.User{}, tx.Error
	}

	return user, nil
}

// 查询所有用户
func SelectAllUsers() ([]model.User, error) {
	var users []model.User

	tx := Db.Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return users, nil
}

// 根据角色ID查询用户
func SelectUsersByRole(r model.UserRole) ([]model.User, error) {
	var users []model.User

	tx := Db.Where("role = ?", r).Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return users, nil
}

// 查询用户头像
func SelectUserAvatarById(id uint64) (string, error) {
	var user model.User
	tx := Db.Where("id = ?", id).First(&user)
	if tx.Error != nil {
		return "", tx.Error
	}

	return user.Avatar, nil
}

// 插入用户
func InsertUser(u model.User) (uint64, error) {
	// 预处理
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	err := u.HashPassword()
	if err != nil {
		return 0, err
	}

	time := time.Now()
	u.CreateTime = time
	u.UpdateTime = time
	tx := Db.Create(&u)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return u.Id, nil
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

	updateTime := time.Now()
	u.CreateTime = updateTime
	u.UpdateTime = updateTime
	tx := Db.Create(&u)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return u.Id, nil
}

// 根据ID更新用户
func UpdateUserById(u model.User) error {
	// 预处理
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	err := u.HashPassword()
	if err != nil {
		return err
	}

	tx := Db.Save(&u)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// 根据ID更新用户（除了密码）
func UpdateUserByIdExceptPassword(u model.User) error {
	// 预处理
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	updateTime := time.Now()
	tx := Db.Model(&model.User{}).Where("id = ?", u.Id).Updates(map[string]interface{}{
		"username":    u.Username,
		"email":       u.Email,
		"signature":   u.Signature,
		"update_time": updateTime,
	})
	if tx.Error != nil {
		return tx.Error
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

	updateTime := time.Now()
	tx := Db.Model(&model.User{}).Where("id = ?", u.Id).Updates(map[string]interface{}{
		"password":    u.Password,
		"update_time": updateTime,
	})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// 根据ID更新用户角色
func UpdateUserRoleById(u model.User) error {
	updateTime := time.Now()
	tx := Db.Model(&model.User{}).Where("id = ?", u.Id).Updates(map[string]interface{}{
		"role":        u.Role,
		"update_time": updateTime,
	})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// 更新用户头像
func UpdateUserAvatarById(id uint64, avatarUrl string) error {
	updateTime := time.Now()
	tx := Db.Model(&model.User{}).Where("id = ?", id).Updates(map[string]interface{}{
		"avatar":      avatarUrl,
		"update_time": updateTime,
	})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// 根据ID删除用户
func DeleteUserById(id uint64) error {
	tx := Db.Where("id = ?", id).Delete(&model.User{})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// 根据邮箱验证密码
func VerifyUserByEmail(u model.User) (uint64, error) {
	password := u.Password

	// 查询用户
	tx := Db.Where("email = ?", u.Email).First(&u)
	if tx.Error != nil {
		return 0, tx.Error
	}

	// 验证密码
	err := u.VerifyByPassword(password)
	if err != nil {
		return 0, err
	}

	return u.Id, nil
}

// 根据ID验证密码
func VerifyUserById(u model.User) (uint64, error) {
	password := u.Password

	// 查询用户
	tx := Db.Where("id = ?", u.Id).First(&u)
	if tx.Error != nil {
		return 0, tx.Error
	}

	// 验证密码
	err := u.VerifyByPassword(password)
	if err != nil {
		return 0, err
	}

	return u.Id, nil
}
