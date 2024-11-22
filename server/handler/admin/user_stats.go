package admin

import (
	"STUOJ/internal/model"
	"STUOJ/internal/service/user"
	"STUOJ/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取用户统计信息
func AdminStatisticsUser(c *gin.Context) {
	// 获取用户统计信息
	stats, err := user.GetStatistics()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", stats))
}

// 获取用户角色统计信息
func AdminStatisticsUserOfRole(c *gin.Context) {
	p, err := utils.GetPeriod(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.RespError(err.Error(), nil))
		return
	}

	// 获取用户统计信息
	stats, err := user.GetStatisticsOfRole(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", stats))
}

// 获取用户注册统计信息
func AdminStatisticsUserOfRegister(c *gin.Context) {
	p, err := utils.GetPeriod(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.RespError(err.Error(), nil))
		return
	}

	// 获取用户统计信息
	stats, err := user.GetStatisticsOfRegister(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", stats))
}
