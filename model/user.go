package model

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

// 角色：0 封禁，1 普通用户，2 管理员，3 超级管理员
type Role uint8

const (
	RoleBanned Role = 0
	RoleUser   Role = 1
	RoleAdmin  Role = 2
	RoleRoot   Role = 3
)

func (r Role) String() string {
	switch r {
	case RoleBanned:
		return "被封禁"
	case RoleUser:
		return "普通用户"
	case RoleAdmin:
		return "管理员"
	case RoleRoot:
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
	Role       Role      `json:"role,omitempty"`
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
