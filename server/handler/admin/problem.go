package admin

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/model"
	"STUOJ/internal/service/problem"
	"STUOJ/internal/service/tag"
	"STUOJ/utils"
	"STUOJ/utils/fps"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
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
	problem, err := problem.SelectById(pid)
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
	testcases, err := dao.SelectTestcasesByProblemId(pid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "获取评测点数据失败",
			Data: nil,
		})
		return
	}

	// 获取题目标签
	tags, err := dao.SelectTagsByProblemId(pid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "获取题目标签失败",
			Data: nil,
		})
		return
	}

	// 获取题解
	solutions, err := dao.SelectSolutionsByProblemId(pid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "获取题解失败",
			Data: nil,
		})
		return
	}

	problemInfo := model.ProblemInfo{
		Problem:   problem,
		Tags:      tags,
		Testcases: testcases,
		Solutions: solutions,
	}

	c.JSON(http.StatusOK, model.Response{
		Code: model.ResponseCodeOk,
		Msg:  "OK",
		Data: problemInfo,
	})
}

// 获取题目列表
func AdminProblemList(c *gin.Context) {
	problems, err := problem.SelectAll()
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

// 根据状态获取题目列表
func AdminProblemListByStatus(c *gin.Context) {
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

	s := model.ProblemStatus(id)
	problems, err := problem.SelectByStatus(s)
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
	p.Id, err = problem.Insert(p)
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
	_, err = dao.InsertProblemHistory(p, uid, model.OperationAdd)
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
	p, err := problem.SelectById(req.Id)
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

	err = problem.UpdateProblemById(p)
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
	_, err = dao.InsertProblemHistory(p, uid, model.OperationUpdate)
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
	_, err = problem.SelectById(pid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "删除失败，题目不存在",
			Data: nil,
		})
		return
	}

	err = problem.DeleteProblemById(pid)
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
	_, err = dao.InsertProblemHistory(p, uid, model.OperationDelete)
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
	phs, err := dao.SelectProblemHistoriesByProblemId(pid)
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
	_, err = problem.SelectById(req.ProblemId)
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
	_, err = tag.SelectById(req.TagId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "添加失败，标签不存在",
			Data: nil,
		})
		return
	}

	// 检查题目标签关系是否存在
	count, err := dao.CountProblemTagByProblemIdAndTagId(req.ProblemId, req.TagId)
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
	err = dao.InsertProblemTag(req.ProblemId, req.TagId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "添加失败",
			Data: nil,
		})
		return
	}

	// 更新题目更新时间
	err = problem.UpdateUpdateTimeById(req.ProblemId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "添加成功，但更新题目更新时间失败",
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

// 删除题目的某个标签
type ReqProblemRemoveTag struct {
	ProblemId uint64 `json:"problem_id,omitempty" binding:"required"`
	TagId     uint64 `json:"tag_id,omitempty" binding:"required"`
}

func AdminProblemRemoveTag(c *gin.Context) {
	var req ReqProblemRemoveTag

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
	_, err = problem.SelectById(req.ProblemId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "删除失败，题目不存在",
			Data: nil,
		})
		return
	}

	// 读取标签
	_, err = tag.SelectById(req.TagId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "删除失败，标签不存在",
			Data: nil,
		})
		return
	}

	// 检查题目标签关系是否存在
	count, err := dao.CountProblemTagByProblemIdAndTagId(req.ProblemId, req.TagId)
	if err != nil || count == 0 {
		if err != nil {
			log.Println(err)
		}
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "删除失败，该题目不存在该标签",
			Data: nil,
		})
		return
	}

	// 初始化标签
	err = dao.DeleteProblemTagByProblemIdAndTagId(req.ProblemId, req.TagId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "删除失败",
			Data: nil,
		})
		return
	}

	// 更新题目更新时间
	err = problem.UpdateUpdateTimeById(req.ProblemId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "删除成功，但更新题目更新时间失败",
			Data: nil,
		})
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.Response{
		Code: model.ResponseCodeOk,
		Msg:  "删除成功",
		Data: nil,
	})
}

func AdminProblemParseFromFps(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "文件上传失败", Data: err})
		return
	}
	dst := fmt.Sprintf("tmp/%s", utils.GetRandKey())
	if err := c.SaveUploadedFile(file, dst); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "文件上传失败", Data: err})
		return
	}
	defer os.Remove(dst)
	f, err := fps.Read(dst)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "文件解析失败", Data: err})
		return
	}
	p := fps.Parse(f)
	c.JSON(http.StatusOK, model.Response{Code: 1, Msg: "文件解析成功", Data: p})
}
