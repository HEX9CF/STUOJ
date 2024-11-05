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
	ID         uint64    `gorm:"primaryKey;autoIncrement;comment:用户ID" json:"id,omitempty"`
	Username   string    `gorm:"type:varchar(255);not null;unique;comment:用户名" json:"username,omitempty"`
	Password   string    `gorm:"type:varchar(255);not null;default:'123456';comment:密码" json:"password,omitempty"`
	Role       UserRole  `gorm:"not null;default:1;comment:角色" json:"role,omitempty"`
	Email      string    `gorm:"type:varchar(255);not null;unique;comment:邮箱" json:"email,omitempty"`
	Avatar     string    `gorm:"type:text;not null;comment:头像URL" json:"avatar,omitempty"`
	Signature  string    `gorm:"type:text;not null;comment:个性签名" json:"signature,omitempty"`
	CreateTime time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time,omitempty"`
	UpdateTime time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"update_time,omitempty"`
}

func (User) TableName() string {
	return "tbl_user"
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
