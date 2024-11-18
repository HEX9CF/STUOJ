package admin

import (
	"STUOJ/internal/model"
	"STUOJ/internal/service/judge"
	"STUOJ/internal/service/user"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// 获取评测机统计信息
func AdminStatisticsJudge(c *gin.Context) {
	statistics, err := judge.GetStatistics()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", statistics))
}

// 获取用户统计信息
type ReqUserStatistics struct {
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

func AdminStatisticsUser(c *gin.Context) {
	var req ReqUserStatistics
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	// 解析时间
	layout := "2006-01-02 15:04:05"
	startTime, err := time.Parse(layout, req.StartTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.RespError("开始时间格式错误", nil))
		return
	}
	endTime, err := time.Parse(layout, req.EndTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.RespError("结束时间格式错误", nil))
		return
	}

	stats, err := user.GetStatistics(startTime, endTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", stats))
}
