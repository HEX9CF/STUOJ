package dao

import (
	"STUOJ/internal/db"
	"STUOJ/internal/model"
)

// 插入用户
func InsertUser(u model.User) (uint64, error) {
	tx := db.Db.Create(&u)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return u.Id, nil
}

// 根据ID查询用户
func SelectUserById(id uint64) (model.User, error) {
	var user model.User

	tx := db.Db.Where("id = ?", id).First(&user)
	if tx.Error != nil {
		return model.User{}, tx.Error
	}

	return user, nil
}

// 根据邮箱查询用户
func SelectUserByEmail(e string) (model.User, error) {
	var user model.User

	tx := db.Db.Where("email = ?", e).First(&user)
	if tx.Error != nil {
		return model.User{}, tx.Error
	}

	return user, nil
}

// 查询所有用户
func SelectAllUsers() ([]model.User, error) {
	var users []model.User

	tx := db.Db.Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return users, nil
}

// 根据角色ID查询用户
func SelectUsersByRole(r model.UserRole) ([]model.User, error) {
	var users []model.User

	tx := db.Db.Where("role = ?", r).Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return users, nil
}

// 根据ID更新用户
func UpdateUserById(u model.User) error {
	tx := db.Db.Save(&u)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// 根据ID删除用户
func DeleteUserById(id uint64) error {
	tx := db.Db.Where("id = ?", id).Delete(&model.User{})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
