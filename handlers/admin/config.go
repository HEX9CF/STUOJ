package admin

import (
	"STUOJ/conf"
	"STUOJ/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取设置列表
func AdminConfigList(c *gin.Context) {
	c.JSON(http.StatusOK, model.Response{
		Code: model.ResponseCodeOk,
		Msg:  "OK",
		Data: conf.Conf,
	})
}
