package judge

import (
	"STUOJ/db"
	"STUOJ/judge"
	"STUOJ/model"
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
	submission.Id, err = db.InsertSubmission(submission)
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

	submission.Status = model.SubmitStatusAC

	chJudgement := make(chan model.Judgement)

	// 提交评测点
	for _, point := range points {
		// 异步评测
		go asyncJudgeSubmit(req, problem, submission, point, chJudgement)
	}

	for _, _ = range points {
		// 接收评测点结果
		judgement := <-chJudgement
		log.Println(judgement)

		// 更新提交数据
		submission.Time = math.Max(submission.Time, judgement.Time)
		submission.Memory = uint64(math.Max(float64(submission.Memory), float64(judgement.Memory)))
		// 如果评测点结果不是AC，更新提交状态
		if judgement.Status != model.SubmitStatusAC {
			submission.Status = judgement.Status
		}
	}

	// 更新提交信息
	err = db.UpdateSubmissionById(submission)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: 0,
			Msg:  "评测失败",
			Data: nil,
		})
		return
	}

	// 返回提交ID
	c.JSON(http.StatusOK, model.Response{
		Code: 1,
		Msg:  "提交成功，返回提交ID",
		Data: submission.Id,
	})
}

func asyncJudgeSubmit(req ReqJudgeSubmit, problem model.Problem, submission model.Submission, point model.TestPoint, c chan model.Judgement) {
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
		TestPointId:   point.Id,
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
	_, err = db.InsertJudgement(judgement)
	if err != nil {
		log.Println(err)
		return
	}

	// 发送评测点结果到通道
	c <- judgement
}
