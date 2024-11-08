package admin

import (
	"STUOJ/external/judge"
	"STUOJ/internal/conf"
	model2 "STUOJ/internal/model"
	"STUOJ/server/model"
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
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "获取配置信息失败",
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code: model.ResponseCodeOk,
		Msg:  "OK",
		Data: configuration,
	})
}
