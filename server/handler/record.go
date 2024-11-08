package handler

import (
	"STUOJ/internal/db/dao"
	model2 "STUOJ/internal/model"
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
		c.JSON(http.StatusBadRequest, model2.Response{
			Code: model2.ResponseCodeError,
			Msg:  "参数错误",
			Data: nil,
		})
		return
	}

	// 获取提交信息
	sid := uint64(id)
	submission, err := dao.SelectSubmissionById(sid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model2.Response{
			Code: model2.ResponseCodeError,
			Msg:  "获取提交信息失败",
			Data: nil,
		})
		return
	}

	// 获取评测结果
	judgements, err := dao.SelectJudgementsBySubmissionId(sid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model2.Response{
			Code: model2.ResponseCodeError,
			Msg:  "获取评测结果失败",
			Data: nil,
		})
		return
	}

	record := model2.Record{
		Submission: submission,
		Judgements: judgements,
	}

	c.JSON(http.StatusOK, model2.Response{
		Code: model2.ResponseCodeOk,
		Msg:  "OK",
		Data: record,
	})
}

// 获取提交记录列表
func RecordList(c *gin.Context) {
	submissions, err := dao.SelectAllSubmissions()
	if err != nil || submissions == nil {
		if err != nil {
			log.Println(err)
		}
		c.JSON(http.StatusOK, model2.Response{
			Code: model2.ResponseCodeError,
			Msg:  "获取失败",
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, model2.Response{
		Code: model2.ResponseCodeOk,
		Msg:  "OK",
		Data: submissions,
	})
}

// 获取题目的提交记录列表
func RecordListOfProblem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model2.Response{
			Code: model2.ResponseCodeError,
			Msg:  "参数错误",
			Data: nil,
		})
		return
	}

	pid := uint64(id)
	submissions, err := dao.SelectSubmissionsByProblemId(pid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model2.Response{
			Code: model2.ResponseCodeError,
			Msg:  "获取提交记录信息失败",
			Data: nil,
		})
		return
	}
	//log.Println(submissions)

	c.JSON(http.StatusOK, model2.Response{
		Code: model2.ResponseCodeOk,
		Msg:  "OK",
		Data: submissions,
	})
}

// 获取用户的提交记录列表
func RecordListOfUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model2.Response{
			Code: model2.ResponseCodeError,
			Msg:  "参数错误",
			Data: nil,
		})
		return
	}

	uid := uint64(id)
	submisssions, err := dao.SelectSubmissionsByUserId(uid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model2.Response{
			Code: model2.ResponseCodeError,
			Msg:  "获取提交记录信息失败",
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, model2.Response{
		Code: model2.ResponseCodeOk,
		Msg:  "OK",
		Data: submisssions,
	})
}
