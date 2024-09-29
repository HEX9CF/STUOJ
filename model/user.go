package model

type Role uint8

const (
	RoleBanned  Role = 1
	RoleUser    Role = 2
	RoleAdmin   Role = 3
	RoleRoot    Role = 4
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

type LoginUserReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}