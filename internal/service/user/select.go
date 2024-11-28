package user

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"errors"
	"log"
)

// 根据ID查询用户
func SelectById(id uint64) (entity.User, error) {
	u, err := dao.SelectUserById(id)
	if err != nil {
		log.Println(err)
		return entity.User{}, errors.New("用户不存在")
	}

	// 不返回密码
	u.Password = ""

	return u, nil
}

// 查询所有用户
func SelectAll() ([]entity.User, error) {
	users, err := dao.SelectAllUsers()
	if err != nil {
		log.Println(err)
		return nil, errors.New("查询用户失败")
	}

	hidePassword(users)

	return users, nil
}

// 根据角色ID查询用户
func SelectByRole(r entity.Role) ([]entity.User, error) {
	users, err := dao.SelectUsersByRole(r)
	if err != nil {
		return nil, errors.New("查询用户失败")
	}

	hidePassword(users)

	return users, nil
}

// 不返回密码
func hidePassword(users []entity.User) {
	for i := range users {
		users[i].Password = ""
	}
}
