package admin

import (
	"STUOJ/external/judge"
	model "STUOJ/internal/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func AdminStatisticsList(c *gin.Context) {
	var err error
	statistics := model.Statistics{}

	statistics.JudgeStatistics, err = judge.GetStatistics()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "获取评测机统计信息失败",
			Data: nil,
		})
		return
	}

	statistics.JudgeSystemInfo, err = judge.GetSystemInfo()
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
