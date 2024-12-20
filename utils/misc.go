package utils

import (
	"STUOJ/internal/entity"
	"encoding/json"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/rand"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func PrettyStruct(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}

func GetRandKey() string {
	rand.Seed(uint64(time.Now().UnixNano()))
	key := make([]rune, 6)
	for i := range key {
		key[i] = letters[rand.Intn(len(letters))]
	}
	return string(key)
}

func IsFileExists(filePath string) (bool, error) {
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil // 文件不存在，返回false和不为nil的error
		}
		return false, err // 其他错误，返回false和错误
	}
	return true, nil // 文件存在，返回true和nil的error
}

func GetUserInfo(c *gin.Context) (entity.Role, uint64) {
	role, exist := c.Get("role")
	if !exist {
		role = entity.RoleVisitor
	}
	id, exist := c.Get("id")
	if !exist {
		id = uint64(0)
	}

	return role.(entity.Role), id.(uint64)
}
