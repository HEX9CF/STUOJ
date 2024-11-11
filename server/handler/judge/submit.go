package judge

import (
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"STUOJ/internal/service/judge"
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
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	// 获取用户ID
	uid, err := utils.GetTokenUid(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, model.RespError("用户ID获取失败", nil))
		return
	}

	// 初始化提交对象
	s := entity.Submission{
		UserId:     uid,
		ProblemId:  req.ProblemId,
		LanguageId: req.LanguageId,
		SourceCode: req.SourceCode,
	}

	// 提交代码
	s.Id, err = judge.Submit(s)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	// 返回提交ID
	c.JSON(http.StatusOK, model.RespOk("提交成功，返回记录提交ID", s.Id))

}

type ReqJudgeTestRun struct {
	LanguageId uint64 `json:"language_id" binding:"required"`
	SourceCode string `json:"source_code" binding:"required"`
	Stdin      string `json:"stdin" binding:"required"`
}

func JudgeTestRun(c *gin.Context) {
	var req ReqJudgeTestRun
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	// 初始化提交对象
	s := entity.Submission{
		LanguageId: req.LanguageId,
		SourceCode: req.SourceCode,
	}

	// 测试运行
	j, err := judge.TestRun(s, req.Stdin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", j))
}
