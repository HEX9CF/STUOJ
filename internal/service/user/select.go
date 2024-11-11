package user

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
)

// 根据ID查询用户
func SelectById(id uint64) (entity.User, error) {
	var user entity.User

	user, err := dao.SelectUserById(id)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

// 查询所有用户
func SelectAll() ([]entity.User, error) {
	var users []entity.User

	users, err := dao.SelectAllUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

// 根据角色ID查询用户
func SelectByRole(r entity.UserRole) ([]entity.User, error) {
	var users []entity.User

	users, err := dao.SelectUsersByRole(r)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// 查询用户头像
func SelectAvatarById(id uint64) (string, error) {
	var user entity.User

	user, err := dao.SelectUserById(id)
	if err != nil {
		return "", err
	}

	return user.Avatar, nil
}

// 根据邮箱验证密码
func VerifyByEmail(u entity.User) (uint64, error) {
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
func VerifyById(u entity.User) (uint64, error) {
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
