package handler

import (
	"STUOJ/internal/model"
	"STUOJ/internal/service/record"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// 获取提交记录信息（提交信息+评测结果）
func RecordInfo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	sid := uint64(id)
	r, err := record.SelectBySubmissionId(sid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
	}

	c.JSON(http.StatusOK, model.RespOk("OK", r))
}

// 获取提交记录列表
func RecordList(c *gin.Context) {
	records, err := record.SelectAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
	}

	c.JSON(http.StatusOK, model.RespOk("OK", records))
}

// 获取题目的提交记录列表
func RecordListOfProblem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	pid := uint64(id)
	records, err := record.SelectByProblemId(pid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", records))
}

// 获取用户的提交记录列表
func RecordListOfUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	uid := uint64(id)
	records, err := record.SelectByUserId(uid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", records))
}
