package handlers

import (
	"STUOJ/db"
	"STUOJ/lskypro"
	"STUOJ/model"
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
	err = db.InsertUser(u)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: 0,
			Msg:  "注册失败，用户名或邮箱已存在",
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

	u.Id, err = db.VerifyUserByEmail(u)
	if err != nil || u.Id == 0 {
		if err != nil {
			log.Println(err)
		}
		c.JSON(http.StatusBadRequest, model.Response{
			Code: 0,
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
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: 0,
			Msg:  "登录失败，生成token失败",
			Data: nil,
		})
		return
	}

	// 登录成功，返回token
	c.JSON(http.StatusOK, model.Response{
		Code: 1,
		Msg:  "登录成功",
		Data: token,
	})
}

func UpdateUserAvatar(c *gin.Context) {
	uploadData, err := lskypro.Upload(c, model.RoleAvatar)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code: 0,
			Msg:  "上传失败",
			Data: nil,
		})
	}
	id, err := utils.ExtractTokenUid(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code: 0,
			Msg:  "获取用户id失败",
			Data: nil,
		})
	}
	err = db.UpdateUserAvatar(id, uploadData.Links.Url)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code: 0,
			Msg:  "更新用户头像失败",
			Data: nil,
		})
	}
	c.JSON(http.StatusOK, model.Response{
		Code: 1,
		Msg:  "更新成功",
		Data: uploadData.Links.Url,
	})
}

func UserAvatar(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code: 0,
			Msg:  "获取用户id失败",
			Data: nil,
		})
	}
	uid := uint64(id)
	avatar, err := db.QueryUserAvatar(uid)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code: 0,
			Msg:  "获取用户头像失败",
			Data: nil,
		})
	}
	c.JSON(http.StatusOK, model.Response{
		Code: 1,
		Msg:  "获取成功",
		Data: avatar,
	})
}

func ThisUserAvatar(c *gin.Context) {
	id, err := utils.ExtractTokenUid(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code: 0,
			Msg:  "获取用户id失败",
			Data: nil,
		})
	}
	avatar, err := db.QueryUserAvatar(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code: 0,
			Msg:  "获取用户头像失败",
			Data: nil,
		})
	}
	c.JSON(http.StatusOK, model.Response{
		Code: 1,
		Msg:  "获取成功",
		Data: avatar,
	})
}
