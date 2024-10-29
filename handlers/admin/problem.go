package admin

import (
	"STUOJ/database"
	"STUOJ/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// 获取题目信息
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
func AdminProblemList(c *gin.Context) {
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
	p.Id, err = database.InsertProblem(p)
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

	// 初始化题目
	p := model.Problem{
		Id:           req.Id,
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
	err = database.UpdateProblemById(p)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "修改失败",
			Data: nil,
		})
		return
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
	_, err = database.SelectProblemById(pid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "删除失败，题目不存在",
			Data: nil,
		})
		return
	}

	err = database.DeleteProblemById(pid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "删除失败",
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