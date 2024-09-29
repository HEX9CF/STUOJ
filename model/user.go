package model

import "time"

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
		return "banned"
	case RoleUser:
		return "user"
	case RoleAdmin:
		return "admin"
	case RoleRoot:
		return "root"
	default:
		return "unknown"
	}
}

type User struct {
	Id         uint64    `json:"id"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Role       Role      `json:"role"`
	Email      string    `json:"email"`
	Avatar     string    `json:"avatar"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

type LoginUserReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
