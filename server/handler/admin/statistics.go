package admin

import (
	"STUOJ/external/judge0"
	"STUOJ/internal/model"
	"STUOJ/internal/service/user"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func AdminStatisticsList(c *gin.Context) {
	var err error
	statistics := model.Statistics{}

	// 获取评测机统计信息
	statistics.JudgeStatistics, err = judge0.GetStatistics()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.RespError("获取评测机统计信息失败", nil))
		return
	}

	// 获取评测机系统信息
	statistics.JudgeSystemInfo, err = judge0.GetSystemInfo()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.RespError("获取评测机系统信息失败", nil))
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, model.RespOk("OK", statistics))
}

func AdminStatisticsUser(c *gin.Context) {
	stats, err := user.GetStatistics()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError("获取用户统计信息失败", nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", stats))
}
