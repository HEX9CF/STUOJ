package admin

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"STUOJ/internal/service/testcase"
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
		c.JSON(http.StatusBadRequest, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "参数错误",
			Data: nil,
		})
		return
	}

	// 获取评测点数据
	tid := uint64(id)
	testcase, err := testcase.SelectById(tid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "获取评测点数据失败",
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code: model.ResponseCodeOk,
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
		c.JSON(http.StatusBadRequest, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "参数错误",
			Data: nil,
		})
		return
	}

	// 初始化题目
	t := entity.Testcase{
		Serial:     req.Serial,
		ProblemId:  req.ProblemId,
		TestInput:  req.TestInput,
		TestOutput: req.TestOutput,
	}
	t.Id, err = testcase.Insert(t)
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
	err = dao.UpdateProblemUpdateTimeById(t.ProblemId)
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
		c.JSON(http.StatusBadRequest, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "参数错误",
			Data: nil,
		})
		return
	}

	// 读取评测点数据
	t, err := testcase.SelectById(req.Id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
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

	err = testcase.UpdateById(t)
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
	err = dao.UpdateProblemUpdateTimeById(t.ProblemId)
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

// 删除评测点数据
func AdminTestcaseRemove(c *gin.Context) {
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

	tid := uint64(id)
	_, err = testcase.SelectById(tid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "删除失败，题目不存在",
			Data: nil,
		})
		return
	}

	err = testcase.DeleteById(tid)
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

func AdminTestcaseDataMake(c *gin.Context) {
	var t model.CommonTestcaseInput
	if err := c.ShouldBindJSON(&t); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "参数错误",
		})
		return
	}
	tc := t.Unfold()
	c.JSON(http.StatusOK, model.Response{
		Code: model.ResponseCodeOk,
		Msg:  "OK",
		Data: tc.String(),
	})
}
