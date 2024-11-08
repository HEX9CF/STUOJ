package admin

import (
	"STUOJ/internal/db/dao"
	model2 "STUOJ/internal/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// 获取评测点数据
func AdminTestcaseInfo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model2.Response{
			Code: model2.ResponseCodeError,
			Msg:  "参数错误",
			Data: nil,
		})
		return
	}

	// 获取评测点数据
	tid := uint64(id)
	testcase, err := dao.SelectTestcaseById(tid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model2.Response{
			Code: model2.ResponseCodeError,
			Msg:  "获取评测点数据失败",
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, model2.Response{
		Code: model2.ResponseCodeOk,
		Msg:  "OK",
		Data: testcase,
	})
}

// 添加评测点数据
type ReqTestcaseAdd struct {
	Serial     uint64 `json:"serial,omitempty" binding:"required"`
	ProblemId  uint64 `json:"problem_id,omitempty" binding:"required"`
	TestInput  string `json:"test_input,omitempty" binding:"required"`
	TestOutput string `json:"test_output,omitempty" binding:"required"`
}

func AdminTestcaseAdd(c *gin.Context) {
	var req ReqTestcaseAdd

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model2.Response{
			Code: model2.ResponseCodeError,
			Msg:  "参数错误",
			Data: nil,
		})
		return
	}

	// 初始化题目
	t := model2.Testcase{
		Serial:     req.Serial,
		ProblemId:  req.ProblemId,
		TestInput:  req.TestInput,
		TestOutput: req.TestOutput,
	}
	t.Id, err = dao.InsertTestcase(t)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model2.Response{
			Code: model2.ResponseCodeError,
			Msg:  "添加失败",
			Data: nil,
		})
		return
	}

	// 更新题目更新时间
	err = dao.UpdateProblemUpdateTimeById(t.ProblemId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model2.Response{
			Code: model2.ResponseCodeError,
			Msg:  "添加成功，但更新题目更新时间失败",
			Data: nil,
		})
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model2.Response{
		Code: model2.ResponseCodeOk,
		Msg:  "添加成功，返回评测点ID",
		Data: t.Id,
	})
}

// 修改评测点数据
type ReqTestcaseModify struct {
	Id         uint64 `json:"id,omitempty" binding:"required"`
	Serial     uint64 `json:"serial,omitempty" binding:"required"`
	ProblemId  uint64 `json:"problem_id,omitempty" binding:"required"`
	TestInput  string `json:"test_input,omitempty" binding:"required"`
	TestOutput string `json:"test_output,omitempty" binding:"required"`
}

func AdminTestcaseModify(c *gin.Context) {
	var req ReqTestcaseModify

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model2.Response{
			Code: model2.ResponseCodeError,
			Msg:  "参数错误",
			Data: nil,
		})
		return
	}

	// 读取评测点数据
	t, err := dao.SelectTestcaseById(req.Id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model2.Response{
			Code: model2.ResponseCodeError,
			Msg:  "修改失败，评测点不存在",
			Data: nil,
		})
		return
	}

	// 修改评测点数据
	t.Serial = req.Serial
	t.ProblemId = req.ProblemId
	t.TestInput = req.TestInput
	t.TestOutput = req.TestOutput

	err = dao.UpdateTestcaseById(t)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model2.Response{
			Code: model2.ResponseCodeError,
			Msg:  "修改失败",
			Data: nil,
		})
		return
	}

	// 更新题目更新时间
	err = dao.UpdateProblemUpdateTimeById(t.ProblemId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model2.Response{
			Code: model2.ResponseCodeError,
			Msg:  "修改成功，但更新题目更新时间失败",
			Data: nil,
		})
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model2.Response{
		Code: model2.ResponseCodeOk,
		Msg:  "修改成功",
		Data: nil,
	})
}

// 删除评测点数据
func AdminTestcaseRemove(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model2.Response{
			Code: model2.ResponseCodeError,
			Msg:  "参数错误",
			Data: nil,
		})
		return
	}

	tid := uint64(id)
	_, err = dao.SelectTestcaseById(tid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model2.Response{
			Code: model2.ResponseCodeError,
			Msg:  "删除失败，题目不存在",
			Data: nil,
		})
		return
	}

	err = dao.DeleteTestcaseById(tid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model2.Response{
			Code: model2.ResponseCodeError,
			Msg:  "删除失败",
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, model2.Response{
		Code: model2.ResponseCodeOk,
		Msg:  "删除成功",
		Data: nil,
	})
}

func AdminTestcaseDataMake(c *gin.Context) {
	var t model2.CommonTestcaseInput
	if err := c.ShouldBindJSON(&t); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model2.Response{
			Code: model2.ResponseCodeError,
			Msg:  "参数错误",
		})
		return
	}
	tc := t.Unfold()
	c.JSON(http.StatusOK, model2.Response{
		Code: model2.ResponseCodeOk,
		Msg:  "OK",
		Data: tc.String(),
	})
}
