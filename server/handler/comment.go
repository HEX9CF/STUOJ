package handler

import (
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"STUOJ/internal/service/comment"
	"STUOJ/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 根据用户ID获取公开评论列表
func CommentPublicListOfUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	uid := uint64(id)
	comments, err := comment.SelectPublicByUserId(uid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", comments))
}

// 根据博客ID获取公开评论列表
func CommentPublicListOfBlog(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	uid := uint64(id)
	comments, err := comment.SelectPublicByBlogId(uid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", comments))
}

// 发表评论
type ReqCommentAdd struct {
	BlogId  uint64 `json:"blog_id,omitempty" binding:"required"`
	Content string `json:"content,omitempty" binding:"required"`
}

func CommentAdd(c *gin.Context) {
	_, id_ := utils.GetUserInfo(c)
	uid := uint64(id_)
	var req ReqCommentAdd

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespOk("参数错误", nil))
		return
	}

	cmt := entity.Comment{
		UserId:  uid,
		BlogId:  req.BlogId,
		Content: req.Content,
	}

	// 插入评论
	cmt.Id, err = comment.Insert(cmt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespOk(err.Error(), nil))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("发布成功，返回评论ID", cmt.Id))
}

// 删除评论
func CommentRemove(c *gin.Context) {
	_, id_ := utils.GetUserInfo(c)
	uid := uint64(id_)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespOk("参数错误", nil))
		return
	}

	// 删除评论
	cid := uint64(id)
	err = comment.DeleteByIdCheckUserId(cid, uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespOk(err.Error(), nil))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("删除成功", nil))
}
