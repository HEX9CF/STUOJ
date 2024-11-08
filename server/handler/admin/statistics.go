package admin

import (
	"STUOJ/external/judge"
	model2 "STUOJ/internal/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func AdminStatisticsList(c *gin.Context) {
	var err error
	statistics := model2.Statistics{}

	statistics.JudgeStatistics, err = judge.GetStatistics()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model2.Response{
			Code: model2.ResponseCodeError,
			Msg:  "获取评测机统计信息失败",
			Data: nil,
		})
		return
	}

	statistics.JudgeSystemInfo, err = judge.GetSystemInfo()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model2.Response{
			Code: model2.ResponseCodeError,
			Msg:  "获取评测机系统信息失败",
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, model2.Response{
		Code: model2.ResponseCodeOk,
		Msg:  "OK",
		Data: statistics,
	})
}
