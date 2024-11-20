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
	"errors"
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
	startTime, endTime, err := parseTime(req.StartTime, req.EndTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.RespError(err.Error(), nil))
		return
	}

	// 获取用户统计信息
	stats, err := user.GetStatistics(startTime, endTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", stats))
}

// 获取提交记录统计信息
type ReqRecordStatistics struct {
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

func AdminStatisticsRecord(c *gin.Context) {
	var req ReqRecordStatistics
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	// 解析时间
	startTime, endTime, err := parseTime(req.StartTime, req.EndTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.RespError(err.Error(), nil))
		return
	}

	// 获取提交记录统计信息
	stats, err := record.GetStatistics(startTime, endTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", stats))
}

// 获取题目统计信息
type ReqProblemStatistics struct {
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

func AdminStatisticsProblem(c *gin.Context) {
	var req ReqProblemStatistics
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	// 解析时间
	startTime, endTime, err := parseTime(req.StartTime, req.EndTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.RespError(err.Error(), nil))
		return
	}

	// 获取题目统计信息
	stats, err := problem.GetStatistics(startTime, endTime)
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

// 获取博客统计信息
type ReqBlogStatistics struct {
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

func AdminStatisticsBlog(c *gin.Context) {
	var req ReqBlogStatistics
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	// 解析时间
	startTime, endTime, err := parseTime(req.StartTime, req.EndTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.RespError(err.Error(), nil))
		return
	}

	// 获取博客统计信息
	stats, err := blog.GetStatistics(startTime, endTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", stats))
}

// 获取评论统计信息
type ReqCommentStatistics struct {
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

func AdminStatisticsComment(c *gin.Context) {
	var req ReqBlogStatistics
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	// 解析时间
	startTime, endTime, err := parseTime(req.StartTime, req.EndTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.RespError(err.Error(), nil))
		return
	}

	// 获取评论统计信息
	stats, err := comment.GetStatistics(startTime, endTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", stats))
}

// 解析时间
func parseTime(s string, e string) (time.Time, time.Time, error) {
	layout := "2006-01-02 15:04:05"
	startTime, err := time.Parse(layout, s)
	if err != nil {
		log.Println(err)
		return time.Time{}, time.Time{}, errors.New("开始时间格式错误")
	}
	endTime, err := time.Parse(layout, e)
	if err != nil {
		log.Println(err)
		return time.Time{}, time.Time{}, errors.New("结束时间格式错误")
	}

	return startTime, endTime, nil
}
