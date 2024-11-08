package user

import (
	"STUOJ/internal/db/dao"
	"STUOJ/internal/model"
)

// 根据ID查询用户
func SelectById(id uint64) (model.User, error) {
	var user model.User

	user, err := dao.SelectUserById(id)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

// 查询所有用户
func SelectAll() ([]model.User, error) {
	var users []model.User

	users, err := dao.SelectAllUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

// 根据角色ID查询用户
func SelectByRole(r model.UserRole) ([]model.User, error) {
	var users []model.User

	users, err := dao.SelectUsersByRole(r)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// 查询用户头像
func SelectAvatarById(id uint64) (string, error) {
	var user model.User

	user, err := dao.SelectUserById(id)
	if err != nil {
		return "", err
	}

	return user.Avatar, nil
}

// 根据邮箱验证密码
func VerifyByEmail(u model.User) (uint64, error) {
	password := u.Password

	// 查询用户
	u, err := dao.SelectUserByEmail(u.Email)
	if err != nil {
		return 0, err
	}

	// 验证密码
	err = u.VerifyByPassword(password)
	if err != nil {
		return 0, err
	}

	return u.Id, nil
}

// 根据ID验证密码
func VerifyById(u model.User) (uint64, error) {
	password := u.Password

	// 查询用户
	u, err := dao.SelectUserById(u.Id)
	if err != nil {
		return 0, err
	}

	// 验证密码
	err = u.VerifyByPassword(password)
	if err != nil {
		return 0, err
	}

	return u.Id, nil
}
