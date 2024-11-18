package user

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"STUOJ/utils"
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

	return u, nil
}

// 查询所有用户
func SelectAll() ([]entity.User, error) {
	users, err := dao.SelectAllUsers()
	if err != nil {
		log.Println(err)
		return nil, errors.New("查询用户失败")
	}

	return users, nil
}

// 根据角色ID查询用户
func SelectByRole(r entity.UserRole) ([]entity.User, error) {
	users, err := dao.SelectUsersByRole(r)
	if err != nil {
		return nil, errors.New("查询用户失败")
	}

	return users, nil
}

// 根据邮箱验证密码
func VerifyByEmail(u entity.User) (string, error) {
	password := u.Password

	// 查询用户
	u, err := dao.SelectUserByEmail(u.Email)
	if err != nil {
		log.Println(err)
		return "", errors.New("用户不存在")
	}

	// 验证密码
	err = u.VerifyByPassword(password)
	if err != nil {
		log.Println(err)
		return "", errors.New("用户名或密码错误")
	}

	// 生成token
	token, err := utils.GenerateToken(u.Id)
	if err != nil {
		log.Println(err)
		return "", errors.New("生成token失败")
	}

	return token, nil
}

// 统计用户
func GetStatistics() ([]model.CountByDate, error) {
	stats, err := dao.CountUsersGroupByCrateTime()
	if err != nil {
		log.Println(err)
		return nil, errors.New("统计用户失败")
	}

	return stats, nil
}
