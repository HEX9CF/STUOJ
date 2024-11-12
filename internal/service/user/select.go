package user

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"STUOJ/utils"
	"errors"
	"log"
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
func VerifyByEmail(u entity.User) (string, error) {
	password := u.Password

	// 查询用户
	u, err := dao.SelectUserByEmail(u.Email)
	if err != nil {
		return "", errors.New("用户不存在")
	}

	// 验证密码
	err = u.VerifyByPassword(password)
	if err != nil {
		return "", errors.New("用户名或密码错误")
	}

	// 生成token
	token, err := utils.GenerateToken(u.Id)
	if err != nil || token == "" {
		if err != nil {
			log.Println(err)
		}
		return "", errors.New("生成token失败")
	}

	return token, nil
}
