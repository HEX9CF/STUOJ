package model

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

type LoginUserReq struct {
	Id       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
}
