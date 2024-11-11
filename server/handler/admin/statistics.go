package admin

import (
	"STUOJ/external/judge0"
	"STUOJ/internal/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func AdminStatisticsList(c *gin.Context) {
	var err error
	statistics := model.Statistics{}

	statistics.JudgeStatistics, err = judge0.GetStatistics()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "获取评测机统计信息失败",
			Data: nil,
		})
		return
	}

	statistics.JudgeSystemInfo, err = judge0.GetSystemInfo()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "获取评测机系统信息失败",
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code: model.ResponseCodeOk,
		Msg:  "OK",
		Data: statistics,
	})
}
