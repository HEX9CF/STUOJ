package admin

import (
	"STUOJ/db"
	"STUOJ/model"
	"STUOJ/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// 获取题目信息（题目+评测点数据）
func AdminProblemInfo(c *gin.Context) {
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

	// 获取题目信息
	pid := uint64(id)
	problem, err := db.SelectProblemById(pid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "获取题目信息失败",
			Data: nil,
		})
		return
	}

	// 获取评测点数据
	testcases, err := db.SelectTestcasesByProblemId(pid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "获取评测点数据失败",
			Data: nil,
		})
		return
	}

	problemInfo := model.ProblemInfo{
		Problem:   problem,
		Testcases: testcases,
	}

	c.JSON(http.StatusOK, model.Response{
		Code: model.ResponseCodeOk,
		Msg:  "OK",
		Data: problemInfo,
	})
}

// 获取题目列表
func AdminProblemList(c *gin.Context) {
	problems, err := db.SelectAllProblems()
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

// 添加题目
type ReqProblemAdd struct {
	Title        string                  `json:"title" binding:"required"`
	Source       string                  `json:"source" binding:"required"`
	Difficulty   model.ProblemDifficulty `json:"difficulty" binding:"required"`
	TimeLimit    float64                 `json:"time_limit" binding:"required"`
	MemoryLimit  uint64                  `json:"memory_limit" binding:"required"`
	Description  string                  `json:"description" binding:"required"`
	Input        string                  `json:"input" binding:"required"`
	Output       string                  `json:"output" binding:"required"`
	SampleInput  string                  `json:"sample_input" binding:"required"`
	SampleOutput string                  `json:"sample_output" binding:"required"`
	Hint         string                  `json:"hint" binding:"required"`
	Status       model.ProblemStatus     `json:"status" binding:"required"`
}

func AdminProblemAdd(c *gin.Context) {
	var req ReqProblemAdd

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

	// 初始化题目
	p := model.Problem{
		Title:        req.Title,
		Source:       req.Source,
		Difficulty:   req.Difficulty,
		TimeLimit:    req.TimeLimit,
		MemoryLimit:  req.MemoryLimit,
		Description:  req.Description,
		Input:        req.Input,
		Output:       req.Output,
		SampleInput:  req.SampleInput,
		SampleOutput: req.SampleOutput,
		Hint:         req.Hint,
		Status:       req.Status,
	}
	p.Id, err = db.InsertProblem(p)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "添加失败",
			Data: nil,
		})
		return
	}

	// 添加题目历史记录
	uid, err := utils.GetTokenUid(c)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "修改成功，但添加题目历史记录失败",
			Data: p.Id,
		})
		return
	}
	_, err = db.InsertProblemHistory(p, uid, model.OperationAdd)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "添加成功，但添加题目历史记录失败",
			Data: p.Id,
		})
	}

	// 返回结果
	c.JSON(http.StatusOK, model.Response{
		Code: model.ResponseCodeOk,
		Msg:  "添加成功，返回题目ID",
		Data: p.Id,
	})
}

// 修改题目
type ReqProblemModify struct {
	Id           uint64                  `json:"id" binding:"required"`
	Title        string                  `json:"title" binding:"required"`
	Source       string                  `json:"source" binding:"required"`
	Difficulty   model.ProblemDifficulty `json:"difficulty" binding:"required"`
	TimeLimit    float64                 `json:"time_limit" binding:"required"`
	MemoryLimit  uint64                  `json:"memory_limit" binding:"required"`
	Description  string                  `json:"description" binding:"required"`
	Input        string                  `json:"input" binding:"required"`
	Output       string                  `json:"output" binding:"required"`
	SampleInput  string                  `json:"sample_input" binding:"required"`
	SampleOutput string                  `json:"sample_output" binding:"required"`
	Hint         string                  `json:"hint" binding:"required"`
	Status       model.ProblemStatus     `json:"status" binding:"required"`
}

func AdminProblemModify(c *gin.Context) {
	var req ReqProblemModify

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

	// 读取题目
	p, err := db.SelectProblemById(req.Id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "修改失败，题目不存在",
			Data: nil,
		})
		return
	}

	// 修改题目
	p.Title = req.Title
	p.Source = req.Source
	p.Difficulty = req.Difficulty
	p.TimeLimit = req.TimeLimit
	p.MemoryLimit = req.MemoryLimit
	p.Description = req.Description
	p.Input = req.Input
	p.Output = req.Output
	p.SampleInput = req.SampleInput
	p.SampleOutput = req.SampleOutput
	p.Hint = req.Hint
	p.Status = req.Status

	err = db.UpdateProblemById(p)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "修改失败",
			Data: nil,
		})
		return
	}

	// 添加题目历史记录
	uid, err := utils.GetTokenUid(c)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "修改成功，但添加题目历史记录失败",
			Data: p.Id,
		})
		return
	}
	_, err = db.InsertProblemHistory(p, uid, model.OperationUpdate)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "修改成功，但添加题目历史记录失败",
			Data: p.Id,
		})
	}

	// 返回结果
	c.JSON(http.StatusOK, model.Response{
		Code: model.ResponseCodeOk,
		Msg:  "修改成功",
		Data: nil,
	})
}

// 删除题目
func AdminProblemRemove(c *gin.Context) {
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
	_, err = db.SelectProblemById(pid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "删除失败，题目不存在",
			Data: nil,
		})
		return
	}

	err = db.DeleteProblemById(pid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "删除失败",
			Data: nil,
		})
		return
	}

	// 添加题目历史记录
	uid, err := utils.GetTokenUid(c)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "删除成功，但添加题目历史记录失败",
			Data: nil,
		})
		return
	}
	p := model.Problem{
		Id: pid,
	}
	_, err = db.InsertProblemHistory(p, uid, model.OperationDelete)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "删除成功，但添加题目历史记录失败",
			Data: nil,
		})
	}

	c.JSON(http.StatusOK, model.Response{
		Code: model.ResponseCodeOk,
		Msg:  "删除成功",
		Data: nil,
	})
}

func AdminProblemHistoryList(c *gin.Context) {
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

	// 获取题目历史记录
	pid := uint64(id)
	phs, err := db.SelectProblemHistoriesByProblemId(pid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "获取题目历史记录失败",
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code: model.ResponseCodeOk,
		Msg:  "OK",
		Data: phs,
	})
}

// 添加标签到题目
type ReqProblemAddTag struct {
	ProblemId uint64 `json:"problem_id,omitempty" binding:"required"`
	TagId     uint64 `json:"tag_id,omitempty" binding:"required"`
}

func AdminProblemAddTag(c *gin.Context) {
	var req ReqProblemAddTag

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

	// 读取题目
	_, err = db.SelectProblemById(req.ProblemId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "添加失败，题目不存在",
			Data: nil,
		})
		return
	}

	// 读取标签
	_, err = db.SelectTagById(req.TagId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "添加失败，标签不存在",
			Data: nil,
		})
		return
	}

	// 读取题目标签关系
	count, err := db.CountProblemTagByProblemIdAndTagId(req.ProblemId, req.TagId)
	if err != nil || count > 0 {
		if err != nil {
			log.Println(err)
		}
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "添加失败，该题目已存在该标签",
			Data: nil,
		})
		return
	}

	// 初始化标签
	_, err = db.InsertProblemTag(req.ProblemId, req.TagId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "添加失败",
			Data: nil,
		})
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.Response{
		Code: model.ResponseCodeOk,
		Msg:  "添加成功",
		Data: nil,
	})
}
