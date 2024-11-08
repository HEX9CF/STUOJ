package admin

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// 获取提交记录信息（提交信息+评测结果）
func AdminRecordInfo(c *gin.Context) {
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

	// 获取提交信息
	sid := uint64(id)
	submission, err := dao.SelectSubmissionById(sid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "获取提交信息失败",
			Data: nil,
		})
		return
	}

	// 获取评测结果
	judgements, err := dao.SelectJudgementsBySubmissionId(sid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "获取评测结果失败",
			Data: nil,
		})
		return
	}

	record := model.Record{
		Submission: submission,
		Judgements: judgements,
	}

	c.JSON(http.StatusOK, model.Response{
		Code: model.ResponseCodeOk,
		Msg:  "OK",
		Data: record,
	})
}

// 获取提交记录列表
func AdminRecordList(c *gin.Context) {
	submissions, err := dao.SelectAllSubmissions()
	if err != nil || submissions == nil {
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
		Data: submissions,
	})
}

// 删除提交记录（提交信息+评测结果）
func AdminRecordRemove(c *gin.Context) {
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

	sid := uint64(id)
	_, err = dao.SelectSubmissionById(sid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "删除失败，提交记录不存在",
			Data: nil,
		})
		return
	}

	// 删除提交信息
	err = dao.DeleteSubmissionById(sid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "删除提交信息失败",
			Data: nil,
		})
		return
	}

	// 删除评测结果
	err = dao.DeleteJudgementBySubmissionId(sid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "删除评测结果失败",
			Data: nil,
		})
		return
	}

	// 更新提交更新时间
	err = dao.UpdateSubmissionUpdateTimeById(sid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "更新提交记录更新时间失败",
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code: model.ResponseCodeOk,
		Msg:  "删除成功",
		Data: nil,
	})
}
