package user

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/model"
	"html"
	"strings"
	"time"
)

// 根据ID更新用户
func UpdateById(u model.User) error {
	// 预处理
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	err := u.HashPassword()
	if err != nil {
		return err
	}

	// 读取用户
	user, err := SelectById(u.Id)
	if err != nil {
		return err
	}

	// 更新用户
	updateTime := time.Now()
	user.Username = u.Username
	user.Email = u.Email
	user.Password = u.Password
	user.Signature = u.Signature
	user.Role = u.Role
	user.Avatar = u.Avatar
	user.UpdateTime = updateTime

	err = dao.UpdateUserById(user)
	if err != nil {
		return err
	}

	return nil
}

// 根据ID更新用户（除了密码）
func UpdateByIdExceptPassword(u model.User) error {
	// 预处理
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	// 读取用户
	user, err := SelectById(u.Id)
	if err != nil {
		return err
	}

	updateTime := time.Now()
	user.Username = u.Username
	user.Email = u.Email
	user.Signature = u.Signature
	user.Role = u.Role
	user.Avatar = u.Avatar
	user.UpdateTime = updateTime

	// 更新用户
	err = dao.UpdateUserById(user)
	if err != nil {
		return err
	}

	return nil
}

// 根据ID更新用户密码
func UpdatePasswordById(u model.User) error {
	// 预处理
	err := u.HashPassword()
	if err != nil {
		return err
	}

	// 读取用户
	user, err := SelectById(u.Id)
	if err != nil {
		return err
	}

	updateTime := time.Now()
	user.Password = u.Password
	user.UpdateTime = updateTime

	// 更新用户
	err = dao.UpdateUserById(user)
	if err != nil {
		return err
	}

	return nil
}

// 根据ID更新用户角色
func UpdateRoleById(u model.User) error {
	// 读取用户
	user, err := SelectById(u.Id)
	if err != nil {
		return err
	}

	updateTime := time.Now()
	user.Role = u.Role
	user.UpdateTime = updateTime

	// 更新用户
	err = dao.UpdateUserById(user)
	if err != nil {
		return err
	}

	return nil
}

// 更新用户头像
func UpdateAvatarById(u model.User) error {
	// 读取用户
	user, err := SelectById(u.Id)
	if err != nil {
		return err
	}

	updateTime := time.Now()
	user.Avatar = u.Avatar
	user.UpdateTime = updateTime

	// 更新用户
	err = dao.UpdateUserById(user)
	if err != nil {
		return err
	}

	return nil
}
