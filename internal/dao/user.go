package dao

import (
	"STUOJ/internal/db"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"time"
)

// 插入用户
func InsertUser(u entity.User) (uint64, error) {
	tx := db.Db.Create(&u)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return u.Id, nil
}

// 根据ID查询用户
func SelectUserById(id uint64) (entity.User, error) {
	var user entity.User

	tx := db.Db.Where("id = ?", id).First(&user)
	if tx.Error != nil {
		return entity.User{}, tx.Error
	}

	return user, nil
}

// 根据邮箱查询用户
func SelectUserByEmail(e string) (entity.User, error) {
	var user entity.User

	tx := db.Db.Where("email = ?", e).First(&user)
	if tx.Error != nil {
		return entity.User{}, tx.Error
	}

	return user, nil
}

// 查询所有用户
func SelectAllUsers() ([]entity.User, error) {
	var users []entity.User

	tx := db.Db.Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return users, nil
}

// 根据角色ID查询用户
func SelectUsersByRole(r entity.UserRole) ([]entity.User, error) {
	var users []entity.User

	tx := db.Db.Where("role = ?", r).Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return users, nil
}

// 根据ID更新用户
func UpdateUserById(u entity.User) error {
	tx := db.Db.Model(&u).Where("id = ?", u.Id).Updates(u)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// 根据ID删除用户
func DeleteUserById(id uint64) error {
	tx := db.Db.Where("id = ?", id).Delete(&entity.User{})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// 统计用户数量
func CountUsers() (int64, error) {
	var count int64

	tx := db.Db.Model(&entity.User{}).Count(&count)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return count, nil
}

// 根据创建时间统计用户数量
func CountUsersBetweenCreateTime(startTime time.Time, endTime time.Time) ([]model.CountByDate, error) {
	var countByDate []model.CountByDate

	tx := db.Db.Model(&entity.User{}).Where("create_time between ? and ?", startTime, endTime).Select("date(create_time) as date, count(*) as count").Group("date").Scan(&countByDate)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return countByDate, nil
}
