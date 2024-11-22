package admin

import (
	"STUOJ/internal/model"
	"STUOJ/internal/service/blog"
	"STUOJ/internal/service/comment"
	"STUOJ/internal/service/judge"
	"STUOJ/internal/service/problem"
	"STUOJ/internal/service/record"
	"STUOJ/internal/service/tag"
	"STUOJ/internal/service/user"
	"STUOJ/utils"
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

func AdminStatisticsUser(c *gin.Context) {
	p, err := utils.GetPeriod(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.RespError(err.Error(), nil))
		return
	}

	// 获取用户统计信息
	stats, err := user.GetStatistics(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", stats))
}

func AdminStatisticsRecord(c *gin.Context) {
	p, err := utils.GetPeriod(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.RespError(err.Error(), nil))
		return
	}

	// 获取提交记录统计信息
	stats, err := record.GetStatistics(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", stats))
}

func AdminStatisticsProblem(c *gin.Context) {
	p, err := utils.GetPeriod(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.RespError(err.Error(), nil))
		return
	}

	// 获取题目统计信息
	stats, err := problem.GetStatistics(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", stats))
}

func AdminStatisticsTag(c *gin.Context) {
	// 获取题目统计信息
	stats, err := tag.GetStatistics()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", stats))
}

func AdminStatisticsBlog(c *gin.Context) {
	p, err := utils.GetPeriod(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.RespError(err.Error(), nil))
		return
	}

	// 获取博客统计信息
	stats, err := blog.GetStatistics(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", stats))
}

func AdminStatisticsComment(c *gin.Context) {
	p, err := utils.GetPeriod(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.RespError(err.Error(), nil))
		return
	}

	// 获取评论统计信息
	stats, err := comment.GetStatistics(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", stats))
}
