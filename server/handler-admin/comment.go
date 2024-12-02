package handler_admin

import (
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"STUOJ/internal/service/comment"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// 获取评论列表
func AdminCommentList(c *gin.Context) {
	comments, err := comment.SelectAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("OK", comments))
}

// 添加评论
type ReqCommentAdd struct {
	UserId  uint64               `json:"user_id,omitempty" binding:"required"`
	BlogId  uint64               `json:"blog_id,omitempty" binding:"required"`
	Content string               `json:"content,omitempty" binding:"required"`
	Status  entity.CommentStatus `json:"status,omitempty"`
}

func AdminCommentAdd(c *gin.Context) {
	var req ReqCommentAdd

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespOk("参数错误", nil))
		return
	}

	cmt := entity.Comment{
		UserId:  req.UserId,
		BlogId:  req.BlogId,
		Content: req.Content,
		Status:  req.Status,
	}

	// 插入评论
	cmt.Id, err = comment.Insert(cmt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespOk(err.Error(), nil))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("添加成功，返回评论ID", cmt.Id))
}

// 修改评论
type ReqCommentModify struct {
	Id      uint64               `json:"id,omitempty" binding:"required"`
	UserId  uint64               `json:"user_id,omitempty" binding:"required"`
	BlogId  uint64               `json:"blog_id,omitempty" binding:"required"`
	Content string               `json:"content,omitempty" binding:"required"`
	Status  entity.CommentStatus `json:"status,omitempty"`
}

func AdminCommentModify(c *gin.Context) {
	var req ReqCommentModify

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespOk("参数错误", nil))
		return
	}

	cmt := entity.Comment{
		Id:      req.Id,
		UserId:  req.UserId,
		BlogId:  req.BlogId,
		Content: req.Content,
		Status:  req.Status,
	}

	// 修改评论
	err = comment.UpdateById(cmt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespOk(err.Error(), nil))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("修改成功", nil))
}

// 删除评论
func AdminCommentRemove(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespOk("参数错误", nil))
		return
	}

	// 删除评论
	cid := uint64(id)
	err = comment.DeleteById(cid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespOk(err.Error(), nil))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("删除成功", nil))
}
