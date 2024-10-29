package middlewares

import (
	"STUOJ/conf"
	"STUOJ/db"
	"STUOJ/model"
	"STUOJ/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func TokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 验证token
		err := utils.VerifyToken(c)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusUnauthorized, model.Response{
				Code: model.ResponseCodeError,
				Msg:  "用户未登录或token过期，请重新登录",
				Data: nil,
			})
			c.Abort()
			return
		}

		// 获取用户id
		uid, err := utils.ExtractTokenUid(c)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusUnauthorized, model.Response{
				Code: model.ResponseCodeError,
				Msg:  "token解析失败",
				Data: nil,
			})
			c.Abort()
			return
		}

		// 获取用户信息
		user, err := db.SelectUserById(uid)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, model.Response{
				Code: model.ResponseCodeError,
				Msg:  "获取用户信息失败",
				Data: nil,
			})
			c.Abort()
			return
		}

		// 校验用户角色
		switch user.Role {
		case model.UserRoleBanned:
			c.JSON(http.StatusUnauthorized, model.Response{
				Code: model.ResponseCodeError,
				Msg:  "用户已被封禁",
				Data: nil,
			})
			c.Abort()
			return
		default:
			break
		}

		// 自动刷新token
		tokenAutoRefresh(c)

		// 放行
		c.Next()
	}
}

func TokenAuthAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 验证token
		err := utils.VerifyToken(c)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusUnauthorized, model.Response{
				Code: model.ResponseCodeError,
				Msg:  "用户未登录或token过期，请重新登录",
				Data: nil,
			})
			c.Abort()
			return
		}

		// 获取用户id
		uid, err := utils.ExtractTokenUid(c)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusUnauthorized, model.Response{
				Code: model.ResponseCodeError,
				Msg:  "token解析失败",
				Data: nil,
			})
			c.Abort()
			return
		}

		// 获取用户信息
		user, err := db.SelectUserById(uid)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, model.Response{
				Code: model.ResponseCodeError,
				Msg:  "获取用户信息失败",
				Data: nil,
			})
			c.Abort()
			return
		}

		// 校验用户角色
		switch user.Role {
		case model.UserRoleBanned:
			c.JSON(http.StatusUnauthorized, model.Response{
				Code: model.ResponseCodeError,
				Msg:  "用户已被封禁",
				Data: nil,
			})
			c.Abort()
			return
		case model.UserRoleUser:
			c.JSON(http.StatusUnauthorized, model.Response{
				Code: model.ResponseCodeError,
				Msg:  "用户权限不足",
				Data: nil,
			})
			c.Abort()
			return
		default:
			break
		}

		// 自动刷新token
		tokenAutoRefresh(c)

		// 放行
		c.Next()
	}
}

func tokenAutoRefresh(c *gin.Context) {
	config := conf.Conf.Token
	exp, err := utils.ExtractTokenExpire(c)
	if err != nil {
		log.Println(err)
		return
	}

	// 计算token剩余时间
	timeLeft := exp - uint64(time.Now().Unix())
	//log.Println(timeLeft, config.Refresh)
	if timeLeft > config.Refresh {
		return
	}

	// 获取用户id
	uid, err := utils.ExtractTokenUid(c)
	if err != nil {
		log.Println(err)
		return
	}

	// 生成新token
	token, err := utils.GenerateToken(uid)
	if err != nil {
		log.Println(err)
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{
		"code": model.ResponseCodeRetry,
		"msg":  "token已刷新，请重新发送请求",
		"data": token,
	})
	c.Abort()
}
