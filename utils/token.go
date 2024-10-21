package utils

import (
	"STUOJ/conf"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"strings"
	"time"
)

// 生成token
func GenerateToken(id uint64) (string, error) {
	config := conf.Conf.Token
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["uid"] = id
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(config.Expire)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.Secret))
}

// 提取token
func ExtractToken(c *gin.Context) string {
	bearerToken := c.GetHeader("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

// 验证token
func VerifyToken(c *gin.Context) error {
	config := conf.Conf.Token
	tokenString := ExtractToken(c)
	_, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(config.Secret), nil
	})
	if err != nil {
		return err
	}
	return nil
}

// 提取token中的uid
func ExtractTokenUid(c *gin.Context) (uint64, error) {
	config := conf.Conf.Token
	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(config.Secret), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["uid"]), 10, 64)
		if err != nil {
			return 0, err
		}
		return uid, nil
	}
	return 0, nil
}
