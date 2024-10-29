package handlers

import (
	"STUOJ/database"
	"STUOJ/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// 获取题目信息
func ProblemInfo(c *gin.Context) {
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
	problem, err := database.SelectProblemById(pid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "获取题目信息失败",
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code: model.ResponseCodeOk,
		Msg:  "OK",
		Data: problem,
	})
}

// 获取题目列表
func ProblemList(c *gin.Context) {
	problems, err := database.SelectAllProblems()
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
