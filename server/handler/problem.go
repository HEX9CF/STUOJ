package handler

import (
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"STUOJ/internal/service/problem"
	"STUOJ/internal/service/tag"
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
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	pid := uint64(id)
	pd, err := problem.SelectPublicByProblemId(pid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", pd))
}

// 获取公开题目列表
func ProblemPublicList(c *gin.Context) {
	pds, err := problem.SelectPublic()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", pds))
}

// 获取标签列表
func TagList(c *gin.Context) {
	tags, err := tag.SelectAll()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", tags))
}

// 根据标签获取公开题目列表
func ProblemPublicListOfTagId(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	tid := uint64(id)
	pds, err := problem.SelectPublicByTagId(tid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", pds))
}

// 根据难度获取公开题目列表
func ProblemPublicListOfDifficulty(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	d := entity.Difficulty(id)
	pds, err := problem.SelectPublicByDifficulty(d)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", pds))
}

type ReqProblemPublicListByTitle struct {
	Title string `json:"title"`
}

// 根据标题获取公开题目列表
func ProblemPublicListOfTitle(c *gin.Context) {
	var req ReqProblemPublicListByTitle
	err := c.BindJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	pds, err := problem.SelectPublicLikeTitle(req.Title)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", pds))
}
