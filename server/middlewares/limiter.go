package middlewares

import (
	"STUOJ/internal/model"
	"STUOJ/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Limiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := utils.GetTokenUid(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, model.Response{
				Code: model.ResponseCodeError,
				Msg:  "token无效，获取用户信息失败",
				Data: nil,
			})
			c.Abort()
			return
		}
		if !utils.IdLimit(uid) {
			c.JSON(http.StatusBadRequest, model.Response{
				Code: model.ResponseCodeError,
				Msg:  "请求过于频繁，请稍后再试",
				Data: nil,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
