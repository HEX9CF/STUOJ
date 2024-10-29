package middlewares

import (
	"STUOJ/conf"
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
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 0,
				"msg":  "用户未登录或token过期，请重新登录",
				"data": nil,
			})
			c.Abort()
			return
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

	// 刷新token
	uid, err := utils.ExtractTokenUid(c)
	if err != nil {
		log.Println(err)
		return
	}

	token, err := utils.GenerateToken(uid)
	if err != nil {
		log.Println(err)
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{
		"code": 2,
		"msg":  "token已刷新，请重新发送请求",
		"data": token,
	})
	c.Abort()
}
