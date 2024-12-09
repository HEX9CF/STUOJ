package handler_admin

import (
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"STUOJ/internal/service/blog"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
