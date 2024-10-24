package handlers

import (
	"STUOJ/db"
	"STUOJ/judge"
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

	// 获取题目信息
	problem, err := db.SelectProblemById(req.ProblemId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: 0,
			Msg:  "获取题目信息失败",
			Data: nil,
		})
		return
	}

	// 获取评测点
	points, err := db.SelectTestPointsByProblemId(req.ProblemId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: 0,
			Msg:  "获取评测点失败",
			Data: nil,
		})
		return
	}

	// 提交评测点
	for _, point := range points {
		// 初始化评测点评测对象
		judgeSubmission := model.JudgeSubmission{
			SourceCode:     req.SourceCode,
			LanguageId:     req.LanguageId,
			Stdin:          point.TestInput,
			ExpectedOutput: point.TestOutput,
			CPUTimeLimit:   problem.TimeLimit,
			MemoryLimit:    problem.MemoryLimit,
		}
		//log.Println(judgeSubmission)

		// 发送评测点评测请求
		token, err := judge.Submit(judgeSubmission)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, model.Response{
				Code: 0,
				Msg:  "评测失败",
				Data: nil,
			})
			return
		}
		log.Println(token)

		// 查询评测点结果
		result, err := judge.QueryResult(token)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, model.Response{
				Code: 0,
				Msg:  "查询评测结果失败",
				Data: nil,
			})
			return
		}
		log.Println(result)

		// 更新评测点结果
	}

	c.JSON(http.StatusOK, model.Response{
		Code: 1,
		Msg:  "提交成功",
		Data: nil,
	})
}
