package admin

import (
	"STUOJ/external/judge"
	"STUOJ/internal/conf"
	model2 "STUOJ/internal/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 获取设置列表
func AdminConfigList(c *gin.Context) {
	var err error
	configuration := model2.Configuration{}

	configuration.System = *conf.Conf
	configuration.Judge, err = judge.GetConfigInfo()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model2.Response{
			Code: model2.ResponseCodeError,
			Msg:  "获取配置信息失败",
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, model2.Response{
		Code: model2.ResponseCodeOk,
		Msg:  "OK",
		Data: configuration,
	})
}
