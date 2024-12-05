package handler

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"STUOJ/internal/service/problem"
	"STUOJ/internal/service/tag"
	"STUOJ/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 获取题目信息
func ProblemInfo(c *gin.Context) {
	role, _ := utils.GetUserInfo(c)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	pid := uint64(id)
	pd, err := problem.SelectById(pid, role >= entity.RoleAdmin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", pd))
}

// 获取题目列表
func ProblemList(c *gin.Context) {
	role, _ := utils.GetUserInfo(c)
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}
	size, err := strconv.Atoi(c.Query("size"))
	if err != nil {
		size = 10
	}
	condition := dao.ProblemWhere{}

	if c.Query("title") != "" {
		condition.Title.Set(c.Query("title"))
	}
	if c.Query("difficulty") != "" {
		difficulty, err := strconv.Atoi(c.Query("difficulty"))
		if err != nil {
			log.Println(err)
		} else {
			condition.Difficulty.Set(entity.Difficulty(difficulty))
		}
	}
	if c.Query("tag") != "" {
		tag, err := strconv.Atoi(c.Query("tag"))
		if err != nil {
			log.Println(err)
		} else {
			condition.Tag.Set(uint64(tag))
		}
	}
	if c.Query("status") != "" {
		status, err := strconv.Atoi(c.Query("status"))
		if err != nil {
			log.Println(err)
		} else {
			condition.Status.Set(entity.ProblemStatus(status))
		}
	}
	if role < entity.RoleAdmin {
		condition.Status.Set(entity.ProblemStatusPublic)
	}

	pds, err := problem.SelectProblem(condition, uint64(page), uint64(size))
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
