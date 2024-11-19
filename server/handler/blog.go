package handler

import (
	"STUOJ/internal/model"
	"STUOJ/internal/service/blog"
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

type ReqBlogPublicListByTitle struct {
	Title string `json:"title"`
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
