package handler

import (
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"STUOJ/internal/service/blog"
	"STUOJ/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// 获取公开博客信息
func BlogPublicInfo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	bid := uint64(id)
	b, err := blog.SelectPublicById(bid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", b))
}

// 获取公开博客列表
func BlogPublicList(c *gin.Context) {
	blogs, err := blog.SelectPublic()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", blogs))
}

// 根据用户ID获取公开博客列表
func BlogPublicListOfUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	uid := uint64(id)
	pds, err := blog.SelectPublicByUserId(uid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", pds))
}

// 根据题目ID获取公开博客列表
func BlogPublicListOfProblem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	pid := uint64(id)
	pds, err := blog.SelectPublicByProblemId(pid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", pds))
}

// 根据标题获取公开题目列表
type ReqBlogPublicListByTitle struct {
	Title string `json:"title"`
}

func BlogPublicListOfTitle(c *gin.Context) {
	var req ReqBlogPublicListByTitle
	err := c.BindJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	blogs, err := blog.SelectPublicLikeTitle(req.Title)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", blogs))
}

// 保存博客
type ReqBlogSave struct {
	ProblemId uint64 `json:"problem_id,omitempty" binding:"required"`
	Title     string `json:"title,omitempty" binding:"required"`
	Content   string `json:"content,omitempty" binding:"required"`
}

func BlogSave(c *gin.Context) {
	var req ReqBlogSave

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespOk("参数错误", nil))
		return
	}

	// 获取用户ID
	uid, err := utils.GetTokenUid(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError("获取用户ID失败", nil))
		return
	}

	b := entity.Blog{
		UserId:    uid,
		ProblemId: req.ProblemId,
		Title:     req.Title,
		Content:   req.Content,
	}

	// 插入博客
	b.Id, err = blog.SaveDraft(b)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespOk(err.Error(), nil))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("保存成功，需要提交审核，返回博客ID", b.Id))
}

// 编辑博客
type ReqBlogEdit struct {
	Id        uint64 `json:"id,omitempty" binding:"required"`
	ProblemId uint64 `json:"problem_id,omitempty" binding:"required"`
	Title     string `json:"title,omitempty" binding:"required"`
	Content   string `json:"content,omitempty" binding:"required"`
}

func BlogEdit(c *gin.Context) {
	var req ReqBlogEdit

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespOk("参数错误", nil))
		return
	}

	// 获取用户ID
	uid, err := utils.GetTokenUid(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError("获取用户ID失败", nil))
		return
	}

	b := entity.Blog{
		Id:        req.Id,
		UserId:    uid,
		ProblemId: req.ProblemId,
		Title:     req.Title,
		Content:   req.Content,
	}

	// 修改博客
	err = blog.EditByIdCheckUserId(b)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespOk(err.Error(), nil))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("修改成功，需要提交审核", nil))
}

// 提交博客
func BlogSubmit(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	// 获取用户ID
	uid, err := utils.GetTokenUid(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError("获取用户ID失败", nil))
		return
	}

	bid := uint64(id)
	err = blog.SubmitByIdCheckUserId(bid, uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("发布成功，等待管理员审核", nil))
}

// 删除博客
func BlogRemove(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespOk("参数错误", nil))
		return
	}

	// 获取用户ID
	uid, err := utils.GetTokenUid(c)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.RespError("获取用户ID失败", nil))
		return
	}

	// 删除博客
	bid := uint64(id)
	err = blog.DeleteByIdCheckUserId(bid, uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespOk(err.Error(), nil))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("删除成功", nil))
}
