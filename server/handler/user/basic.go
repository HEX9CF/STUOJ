package user

import (
	"STUOJ/internal/db"
	model2 "STUOJ/internal/model"
	"STUOJ/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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
		c.JSON(http.StatusBadRequest, model2.Response{
			Code: model2.ResponseCodeError,
			Msg:  "参数错误",
			Data: nil,
		})
		return
	}

	// 初始化用户
	u := model2.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}
	u.Id, err = db.InsertUserForRegister(u)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model2.Response{
			Code: model2.ResponseCodeError,
			Msg:  "注册失败，用户名或邮箱已存在",
			Data: nil,
		})
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model2.Response{
		Code: model2.ResponseCodeOk,
		Msg:  "注册成功，返回用户ID",
		Data: u.Id,
	})
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
		c.JSON(http.StatusBadRequest, model2.Response{
			Code: model2.ResponseCodeError,
			Msg:  "参数错误",
			Data: nil,
		})
		return
	}

	// 校验参数
	if req.Password == "" || req.Email == "" {
		c.JSON(http.StatusBadRequest, model2.Response{
			Code: model2.ResponseCodeError,
			Msg:  "参数错误，邮箱或密码不能为空",
			Data: nil,
		})
		return
	}

	// 初始化用户
	u := model2.User{
		Email:    req.Email,
		Password: req.Password,
	}

	u.Id, err = db.VerifyUserByEmail(u)
	if err != nil || u.Id == 0 {
		if err != nil {
			log.Println(err)
		}
		c.JSON(http.StatusBadRequest, model2.Response{
			Code: model2.ResponseCodeError,
			Msg:  "登录失败，用户名或密码错误",
			Data: nil,
		})
		return
	}

	// 生成token
	//token := "{test token}"
	token, err := utils.GenerateToken(u.Id)
	if err != nil || token == "" {
		if err != nil {
			log.Println(err)
		}
		c.JSON(http.StatusInternalServerError, model2.Response{
			Code: model2.ResponseCodeError,
			Msg:  "登录失败，生成token失败",
			Data: nil,
		})
		return
	}

	// 登录成功，返回token
	c.JSON(http.StatusOK, model2.Response{
		Code: model2.ResponseCodeOk,
		Msg:  "登录成功，返回token",
		Data: token,
	})
}
