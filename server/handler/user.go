package handler

import (
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"STUOJ/internal/service/user"
	"STUOJ/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// 用户注册
type ReqUserRegister struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func UserRegister(c *gin.Context) {
	var req ReqUserRegister

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	// 初始化用户
	u := entity.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}
	u.Id, err = user.Register(u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.RespOk("注册成功，返回用户ID", u.Id))
}

// 用户登录
type ReqUserLogin struct {
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func UserLogin(c *gin.Context) {
	var req ReqUserLogin

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	// 初始化用户
	u := entity.User{
		Email:    req.Email,
		Password: req.Password,
	}

	token, err := user.VerifyByEmail(u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	// 登录成功，返回token
	c.JSON(http.StatusOK, model.RespOk("登录成功，返回token", token))
}

// 获取用户信息
func UserInfo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	uid := uint64(id)
	u, err := user.SelectById(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", u))
}

// 获取当前用户id
func UserCurrentId(c *gin.Context) {
	id, err := utils.GetTokenUid(c)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, model.RespError("用户未登录", nil))
	}

	c.JSON(http.StatusOK, model.RespOk("OK", id))
}
