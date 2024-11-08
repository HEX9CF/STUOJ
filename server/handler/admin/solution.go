package admin

import (
	"STUOJ/internal/db/dao"
	model2 "STUOJ/internal/model"
	"STUOJ/server/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// 获取题解列表
func AdminSolutionList(c *gin.Context) {
	solutions, err := dao.SelectAllSolutions()
	if err != nil || solutions == nil {
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
		Data: solutions,
	})
}

// 根据题目ID获取题解列表
func AdminSolutionListByProblemId(c *gin.Context) {
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
	solutions, err := dao.SelectSolutionsByProblemId(pid)
	if err != nil || solutions == nil {
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
		Data: solutions,
	})
}

// 获取题解数据
func AdminSolutionInfo(c *gin.Context) {
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

	// 获取评测点数据
	sid := uint64(id)
	solution, err := dao.SelectSolutionById(sid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "获取题解数据失败",
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code: model.ResponseCodeOk,
		Msg:  "OK",
		Data: solution,
	})
}

// 添加题解
type ReqSolutionAdd struct {
	LanguageId uint64 `json:"language_id,omitempty" binding:"required"`
	ProblemId  uint64 `json:"problem_id,omitempty" binding:"required"`
	SourceCode string `json:"source_code,omitempty" binding:"required"`
}

func AdminSolutionAdd(c *gin.Context) {
	var req ReqSolutionAdd

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

	s := model2.Solution{
		LanguageId: req.LanguageId,
		ProblemId:  req.ProblemId,
		SourceCode: req.SourceCode,
	}
	s.Id, err = dao.InsertSolution(s)
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
	err = dao.UpdateProblemUpdateTimeById(s.ProblemId)
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
		Msg:  "添加成功，返回题解ID",
		Data: s.Id,
	})
}

// 修改题解
type ReqSolutionModify struct {
	Id         uint64 `json:"id,omitempty" binding:"required"`
	LanguageId uint64 `json:"language_id,omitempty" binding:"required"`
	ProblemId  uint64 `json:"problem_id,omitempty" binding:"required"`
	SourceCode string `json:"source_code,omitempty" binding:"required"`
}

func AdminSolutionModify(c *gin.Context) {
	var req ReqSolutionModify

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

	// 读取题解
	s, err := dao.SelectSolutionById(req.Id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "修改失败，题解不存在",
			Data: nil,
		})
		return
	}

	// 修改题解
	s.LanguageId = req.LanguageId
	s.ProblemId = req.ProblemId
	s.SourceCode = req.SourceCode

	err = dao.UpdateSolutionById(s)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "修改失败",
			Data: nil,
		})
		return
	}

	// 更新题目更新时间
	err = dao.UpdateProblemUpdateTimeById(s.ProblemId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "修改成功，但更新题目更新时间失败",
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

// 删除题解
func AdminSolutionRemove(c *gin.Context) {
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
	_, err = dao.SelectSolutionById(sid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "删除失败，题解不存在",
			Data: nil,
		})
		return
	}

	err = dao.DeleteSolutionById(sid)
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
