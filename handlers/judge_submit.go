package handlers

import (
	"STUOJ/db"
	"STUOJ/model"
	"STUOJ/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 提交评测
type ReqJudgeSubmit struct {
	LanguageId uint64 `json:"language_id" binding:"required"`
	ProblemId  uint64 `json:"problem_id" binding:"required"`
	SourceCode string `json:"source_code" binding:"required"`
}

func JudgeSubmit(c *gin.Context) {
	var req ReqJudgeSubmit

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.Response{
			Code: 0,
			Msg:  "参数错误",
			Data: nil,
		})
		return
	}

	// 获取用户ID
	uid, err := utils.ExtractTokenUid(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, model.Response{
			Code: 0,
			Msg:  "用户未登录",
			Data: nil,
		})
		return
	}

	// 获取代码长度
	length := uint64(len(req.SourceCode))

	// 初始化提交对象
	submission := model.Submission{
		UserId:     uid,
		ProblemId:  req.ProblemId,
		Status:     0,
		Score:      0,
		Length:     length,
		LanguageId: req.LanguageId,
		SourceCode: req.SourceCode,
	}

	// 插入提交
	err = db.InsertSubmission(submission)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: 0,
			Msg:  "提交失败",
			Data: nil,
		})
		return
	}

	// TODO: 提交评测

	c.JSON(http.StatusOK, model.Response{
		Code: 1,
		Msg:  "提交成功",
		Data: nil,
	})
}
