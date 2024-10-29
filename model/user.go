package model

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

// 角色：0 封禁，1 普通用户，2 管理员，3 超级管理员
type UserRole uint8

const (
	UserRoleBanned UserRole = 0
	UserRoleUser   UserRole = 1
	UserRoleAdmin  UserRole = 2
	UserRoleRoot   UserRole = 3
)

func (r UserRole) String() string {
	switch r {
	case UserRoleBanned:
		return "被封禁"
	case UserRoleUser:
		return "普通用户"
	case UserRoleAdmin:
		return "管理员"
	case UserRoleRoot:
		return "超级管理员"
	default:
		return "未知角色"
	}
}

// 用户
type User struct {
	Id         uint64    `json:"id,omitempty"`
	Username   string    `json:"username,omitempty"`
	Password   string    `json:"password,omitempty"`
	Role       UserRole  `json:"role,omitempty"`
	Email      string    `json:"email,omitempty"`
	Avatar     string    `json:"avatar,omitempty"`
	Signature  string    `json:"signature,omitempty"`
	CreateTime time.Time `json:"create_time,omitempty"`
	UpdateTime time.Time `json:"update_time,omitempty"`
}

// 对密码进行哈希
func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// 验证密码
func (u *User) VerifyByPassword(pw string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pw))
}

// 验证密码
func (u *User) VerifyByHashedPassword(hpw string) error {
	return bcrypt.CompareHashAndPassword([]byte(hpw), []byte(u.Password))
}
