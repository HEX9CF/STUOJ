package handlers

import (
	"STUOJ/db"
	"STUOJ/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

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
		c.JSON(http.StatusBadRequest, model.Response{
			Code: 0,
			Msg:  "参数错误",
			Data: nil,
		})
		return
	}

	// 校验参数
	if req.Username == "" || req.Password == "" || req.Email == "" {
		c.JSON(http.StatusBadRequest, model.Response{
			Code: 0,
			Msg:  "参数错误，用户名、邮箱或密码不能为空",
			Data: nil,
		})
		return
	}

	// 初始化用户
	u := model.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}
	err = db.SaveUser(u)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: 0,
			Msg:  "注册失败",
			Data: nil,
		})
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.Response{
		Code: 1,
		Msg:  "注册成功",
		Data: nil,
	})
}

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
		c.JSON(http.StatusBadRequest, model.Response{
			Code: 0,
			Msg:  "参数错误",
			Data: nil,
		})
		return
	}

	// 校验参数
	if req.Password == "" || req.Email == "" {
		c.JSON(http.StatusBadRequest, model.Response{
			Code: 0,
			Msg:  "参数错误，邮箱或密码不能为空",
			Data: nil,
		})
		return
	}

	// 初始化用户
	u := model.User{
		Email:    req.Email,
		Password: req.Password,
	}

	id, err := db.LoginUserByEmail(u)
	if err != nil || id == 0 {
		if err != nil {
			log.Println(err)
		}
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: 0,
			Msg:  "登录失败，用户名或密码错误",
			Data: nil,
		})
		return
	}

	// 生成token
	token := "test token"

	// 登录成功，返回token
	c.JSON(http.StatusOK, model.Response{
		Code: 1,
		Msg:  "登录成功",
		Data: token,
	})
}

func UserLogout(c *gin.Context) {
}

func UserData(c *gin.Context) {
}
