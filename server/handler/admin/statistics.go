package admin

import (
	"STUOJ/internal/model"
	"STUOJ/internal/service/judge"
	"STUOJ/internal/service/user"
	"github.com/gin-gonic/gin"
	"net/http"
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
func AdminStatisticsUser(c *gin.Context) {
	stats, err := user.GetStatistics()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", stats))
}
