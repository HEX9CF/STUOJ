package admin

import (
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"STUOJ/internal/service/blog"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// 获取博客列表
func AdminBlogList(c *gin.Context) {
	blogs, err := blog.SelectAll()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("OK", blogs))
}

// 获取博客数据
func AdminBlogInfo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespOk("参数错误", nil))
		return
	}

	// 获取博客数据
	sid := uint64(id)
	s, err := blog.SelectById(sid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespOk(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", s))
}

// 添加博客
type ReqBlogAdd struct {
	UserId    uint64            `json:"user_id,omitempty" binding:"required"`
	ProblemId uint64            `json:"problem_id,omitempty" binding:"required"`
	Title     string            `json:"title,omitempty" binding:"required"`
	Content   string            `json:"content,omitempty" binding:"required"`
	Status    entity.BlogStatus `json:"status,omitempty"`
}

func AdminBlogAdd(c *gin.Context) {
	var req ReqBlogAdd

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespOk("参数错误", nil))
		return
	}

	b := entity.Blog{
		UserId:    req.UserId,
		ProblemId: req.ProblemId,
		Title:     req.Title,
		Content:   req.Content,
		Status:    req.Status,
	}

	// 插入博客
	b.Id, err = blog.Insert(b)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespOk(err.Error(), nil))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("添加成功，返回博客ID", b.Id))
}

// 修改博客
type ReqBlogModify struct {
	Id        uint64            `json:"id,omitempty" binding:"required"`
	UserId    uint64            `json:"user_id,omitempty" binding:"required"`
	ProblemId uint64            `json:"problem_id,omitempty" binding:"required"`
	Title     string            `json:"title,omitempty" binding:"required"`
	Content   string            `json:"content,omitempty" binding:"required"`
	Status    entity.BlogStatus `json:"status,omitempty"`
}

func AdminBlogModify(c *gin.Context) {
	var req ReqBlogModify

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespOk("参数错误", nil))
		return
	}

	b := entity.Blog{
		Id:        req.Id,
		UserId:    req.UserId,
		ProblemId: req.ProblemId,
		Title:     req.Title,
		Content:   req.Content,
		Status:    req.Status,
	}

	// 修改博客
	err = blog.UpdateById(b)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespOk(err.Error(), nil))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("修改成功", nil))
}

// 删除博客
func AdminBlogRemove(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespOk("参数错误", nil))
		return
	}

	// 删除博客
	bid := uint64(id)
	err = blog.DeleteById(bid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespOk(err.Error(), nil))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("删除成功", nil))
}
