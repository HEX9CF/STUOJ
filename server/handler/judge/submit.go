package judge

import (
	"STUOJ/external/judge"
	"STUOJ/internal/dao"
	"STUOJ/internal/model"
	"STUOJ/internal/service/problem"
	"STUOJ/internal/service/testcase"
	"STUOJ/utils"
	"github.com/gin-gonic/gin"
	"log"
	"math"
	"net/http"
	"strconv"
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
			Code: model.ResponseCodeError,
			Msg:  "参数错误",
			Data: nil,
		})
		return
	}

	// 获取用户ID
	uid, err := utils.GetTokenUid(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, model.Response{
			Code: model.ResponseCodeError,
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
	submission.Id, err = dao.InsertSubmission(submission)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "提交失败",
			Data: nil,
		})
		return
	}

	// 获取题目信息
	problem, err := problem.SelectById(req.ProblemId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "获取题目信息失败",
			Data: nil,
		})
		return
	}

	// 获取评测点
	testcases, err := testcase.SelectByProblemId(req.ProblemId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "获取评测点数据失败",
			Data: nil,
		})
		return
	}

	// 返回提交ID
	c.JSON(http.StatusOK, model.Response{
		Code: model.ResponseCodeOk,
		Msg:  "提交成功，返回记录提交ID",
		Data: submission.Id,
	})

	submission.Status = model.SubmitStatusAC
	chJudgement := make(chan model.Judgement)

	// 提交评测点
	for _, testcase := range testcases {
		// 异步评测
		go asyncJudgeSubmit(req, problem, submission, testcase, chJudgement)
	}

	for _, _ = range testcases {
		// 接收评测点结果
		judgement := <-chJudgement
		//log.Println(judgement)

		// 更新提交数据
		submission.Time = math.Max(submission.Time, judgement.Time)
		submission.Memory = max(submission.Memory, judgement.Memory)
		// 如果评测点结果不是AC，更新提交状态
		if judgement.Status != model.SubmitStatusAC {
			if submission.Status != model.SubmitStatusWA {
				submission.Status = max(submission.Status, judgement.Status)
			}
		}
	}

	// 更新提交信息
	err = dao.UpdateSubmissionById(submission)
	if err != nil {
		log.Println(err)
		return
	}
}

func asyncJudgeSubmit(req ReqJudgeSubmit, problem model.Problem, submission model.Submission, testcase model.Testcase, c chan model.Judgement) {
	// 初始化评测点评测对象
	judgeSubmission := model.JudgeSubmission{
		SourceCode:     req.SourceCode,
		LanguageId:     req.LanguageId,
		Stdin:          testcase.TestInput,
		ExpectedOutput: testcase.TestOutput,
		CPUTimeLimit:   problem.TimeLimit,
		MemoryLimit:    problem.MemoryLimit,
	}
	//log.Println(judgeSubmission)

	// 发送评测点评测请求
	result, err := judge.Submit(judgeSubmission)
	if err != nil {
		log.Println(err)
		return
	}
	//log.Println(result)

	// 解析时间
	time := float64(0)
	if result.Time != "" {
		time, err = strconv.ParseFloat(result.Time, 64)
		if err != nil {
			log.Println(err)
			return
		}
	}

	// 初始化评测点结果对象
	judgement := model.Judgement{
		SubmissionId:  submission.Id,
		TestcaseId:    testcase.Id,
		Time:          time,
		Memory:        uint64(result.Memory),
		Stdout:        result.Stdout,
		Stderr:        result.Stderr,
		CompileOutput: result.CompileOutput,
		Message:       result.Message,
		Status:        model.SubmitStatus(result.Status.Id),
	}
	//log.Println(judgement)

	// 更新评测点结果
	_, err = dao.InsertJudgement(judgement)
	if err != nil {
		log.Println(err)
		return
	}

	// 更新提交更新时间
	err = dao.UpdateSubmissionUpdateTimeById(submission.Id)
	if err != nil {
		log.Println(err)
		return
	}

	// 发送评测点结果到通道
	c <- judgement
}

type ReqJudgeTestRun struct {
	LanguageId uint64 `json:"language_id" binding:"required"`
	SourceCode string `json:"source_code" binding:"required"`
	Stdin      string `json:"stdin" binding:"required"`
}

func JudgeTestRun(c *gin.Context) {
	var t ReqJudgeTestRun
	err := c.ShouldBindJSON(&t)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "参数错误",
			Data: err,
		})
		return
	}

	s := model.JudgeSubmission{
		SourceCode: t.SourceCode,
		LanguageId: t.LanguageId,
		Stdin:      t.Stdin,
	}

	res, err := judge.Submit(s)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "提交失败",
			Data: err,
		})
		return
	}

	time, err := strconv.ParseFloat(res.Time, 64)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "解析时间失败",
			Data: nil,
		})
	}

	j := model.Judgement{
		Stdout: res.Stdout,
		Time:   time,
		Memory: uint64(res.Memory),
	}

	c.JSON(http.StatusOK, model.Response{
		Code: model.ResponseCodeOk,
		Msg:  "OK",
		Data: j,
	})
}
