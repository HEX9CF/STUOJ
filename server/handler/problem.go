package handler

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/model"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 获取公开题目信息
func ProblemPublicInfo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "参数错误",
			Data: nil,
		})
		return
	}

	pid := uint64(id)
	problem, err := dao.SelectProblemByStatusAndId(pid, model.ProblemStatusPublic)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "获取题目信息失败",
			Data: nil,
		})
		return
	}

	// 获取题目标签
	tags, err := dao.SelectTagsByProblemId(pid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "获取题目标签失败",
			Data: nil,
		})
		return
	}

	// 初始化题目信息
	problemInfo := model.ProblemInfo{
		Problem: problem,
		Tags:    tags,
	}

	c.JSON(http.StatusOK, model.Response{
		Code: model.ResponseCodeOk,
		Msg:  "OK",
		Data: problemInfo,
	})
}

// 获取公开题目列表
func ProblemPublicList(c *gin.Context) {
	problems, err := dao.SelectAllProblemsByStatus(model.ProblemStatusPublic)
	if err != nil || problems == nil {
		if err != nil {
			log.Println(err)
		}
		c.JSON(http.StatusOK, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "获取失败",
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code: model.ResponseCodeOk,
		Msg:  "OK",
		Data: problems,
	})
}

// 获取标签列表
func TagList(c *gin.Context) {
	tags, err := dao.SelectAllTags()
	if err != nil || tags == nil {
		if err != nil {
			log.Println(err)
		}
		c.JSON(http.StatusOK, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "获取失败",
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code: model.ResponseCodeOk,
		Msg:  "OK",
		Data: tags,
	})
}

// 根据标签获取公开题目列表
func ProblemPublicListByTagId(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "参数错误",
			Data: nil,
		})
		return
	}

	tid := uint64(id)
	problems, err := dao.SelectProblemsByTagIdAndStatus(tid, model.ProblemStatusPublic)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "获取失败",
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code: model.ResponseCodeOk,
		Msg:  "OK",
		Data: problems,
	})
}

// 根据难度获取公开题目列表
func ProblemPublicListByDifficulty(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "参数错误",
			Data: nil,
		})
		return
	}

	d := model.ProblemDifficulty(id)
	problems, err := dao.SelectProblemsByDifficultyAndStatus(d, model.ProblemStatusPublic)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "获取失败",
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code: model.ResponseCodeOk,
		Msg:  "OK",
		Data: problems,
	})
}

type ReqProblemPublicListByTitle struct {
	Title string `json:"title"`
}

// 根据标题获取公开题目列表
func ProblemPublicListByTitle(c *gin.Context) {
	var req ReqProblemPublicListByTitle
	err := c.BindJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "参数错误",
			Data: nil,
		})
		return
	}

	problems, err := dao.SelectProblemsLikeTitleByStatus(req.Title, model.ProblemStatusPublic)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "获取失败",
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code: model.ResponseCodeOk,
		Msg:  "OK",
		Data: problems,
	})
}
