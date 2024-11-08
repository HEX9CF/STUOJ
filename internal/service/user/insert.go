package user

import (
	"STUOJ/internal/db/dao"
	"STUOJ/internal/model"
	"html"
	"strings"
	"time"
)

// 插入用户
func InsertUser(u model.User) (uint64, error) {
	// 预处理
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	err := u.HashPassword()
	if err != nil {
		return 0, err
	}

	updateTime := time.Now()
	u.CreateTime = updateTime
	u.UpdateTime = updateTime

	u.Id, err = dao.InsertUser(u)
	if err != nil {
		return 0, err
	}

	return u.Id, nil
}

// 插入用户（注册）
func Register(u model.User) (uint64, error) {
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

	u.Id, err = dao.InsertUser(u)
	if err != nil {
		return 0, err
	}

	return u.Id, nil
}
