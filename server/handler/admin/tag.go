package admin

import (
	"STUOJ/internal/db"
	model2 "STUOJ/internal/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// 获取标签列表
func AdminTagList(c *gin.Context) {
	tags, err := db.SelectAllTags()
	if err != nil || tags == nil {
		if err != nil {
			log.Println(err)
		}
		c.JSON(http.StatusOK, model2.Response{
			Code: model2.ResponseCodeError,
			Msg:  "获取失败",
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, model2.Response{
		Code: model2.ResponseCodeOk,
		Msg:  "OK",
		Data: tags,
	})
}

// 添加标签
type ReqTagAdd struct {
	Name string `json:"name,omitempty" binding:"required"`
}

func AdminTagAdd(c *gin.Context) {
	var req ReqTagAdd

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

	// 初始化标签
	t := model2.Tag{
		Name: req.Name,
	}
	t.Id, err = db.InsertTag(t)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model2.Response{
			Code: model2.ResponseCodeError,
			Msg:  "添加失败，标签不能重复",
			Data: nil,
		})
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model2.Response{
		Code: model2.ResponseCodeOk,
		Msg:  "添加成功，返回标签ID",
		Data: t.Id,
	})
}

// 修改标签数据
type ReqTagModify struct {
	Id   uint64 `json:"id,omitempty" binding:"required"`
	Name string `json:"name,omitempty" binding:"required"`
}

func AdminTagModify(c *gin.Context) {
	var req ReqTagModify

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

	// 读取标签
	t, err := db.SelectTagById(req.Id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model2.Response{
			Code: model2.ResponseCodeError,
			Msg:  "修改失败，标签不存在",
			Data: nil,
		})
		return
	}

	// 修改标签
	t.Name = req.Name

	err = db.UpdateTagById(t)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model2.Response{
			Code: model2.ResponseCodeError,
			Msg:  "修改失败，标签不能重复",
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

// 删除标签
func AdminTagRemove(c *gin.Context) {
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
	_, err = db.SelectTagById(tid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model2.Response{
			Code: model2.ResponseCodeError,
			Msg:  "删除失败，标签不存在",
			Data: nil,
		})
		return
	}

	err = db.DeleteTagById(tid)
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
