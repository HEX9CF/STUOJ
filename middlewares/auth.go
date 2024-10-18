package middlewares

import (
	"STUOJ/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func TokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := utils.VerifyToken(c)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 0,
				"msg":  "用户未登录",
				"data": nil,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
