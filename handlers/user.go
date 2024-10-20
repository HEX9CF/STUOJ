package handlers

import (
	"STUOJ/db"
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

// 获取当前用户id
func UserCurrentId(c *gin.Context) {
	id, err := utils.ExtractTokenUid(c)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, model.Response{
			Code: 0,
			Msg:  "用户未登录",
			Data: nil,
		})
	}

	c.JSON(http.StatusOK, model.Response{
		Code: 1,
		Msg:  "OK",
		Data: id,
	})
}

// 获取用户信息
func UserInfo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.Response{
			Code: 0,
			Msg:  "参数错误",
			Data: nil,
		})
		return
	}

	uid := uint64(id)
	user, err := db.SelectUserById(uid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: 0,
			Msg:  "获取用户信息失败",
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code: 1,
		Msg:  "OK",
		Data: user,
	})
}

// 获取用户列表
func UserList(c *gin.Context) {
	users, err := db.SelectAllUsers()
	if err != nil || users == nil {
		if err != nil {
			log.Println(err)
		}
		c.JSON(http.StatusOK, model.Response{
			Code: 0,
			Msg:  "获取失败",
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code: 1,
		Msg:  "OK",
		Data: users,
	})
}

// 修改用户信息
type ReqUserModify struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func UserModify(c *gin.Context) {
	var req ReqUserModify

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
	if req.Username == "" || req.Email == "" {
		c.JSON(http.StatusBadRequest, model.Response{
			Code: 0,
			Msg:  "参数错误，用户名或邮箱不能为空",
			Data: nil,
		})
		return
	}

	// 获取用户id
	id, err := utils.ExtractTokenUid(c)

	// 初始化用户
	u := model.User{
		Id:       id,
		Username: req.Username,
		Email:    req.Email,
	}
	err = db.UpdateUserById(u)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: 0,
			Msg:  "修改失败, 用户名或邮箱已存在",
			Data: nil,
		})
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.Response{
		Code: 1,
		Msg:  "修改成功",
		Data: nil,
	})
}

// 修改用户密码
type ReqUserChangePassword struct {
	Password string `json:"password" binding:"required"`
}

func UserChangePassword(c *gin.Context) {
	var req ReqUserChangePassword

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
	if req.Password == "" {
		c.JSON(http.StatusBadRequest, model.Response{
			Code: 0,
			Msg:  "参数错误，密码不能为空",
			Data: nil,
		})
		return
	}

	// 获取用户id
	id, err := utils.ExtractTokenUid(c)

	// 初始化用户
	u := model.User{
		Id:       id,
		Password: req.Password,
	}
	err = db.UpdateUserPasswordById(u)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: 0,
			Msg:  "修改失败",
			Data: nil,
		})
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.Response{
		Code: 1,
		Msg:  "修改成功",
		Data: nil,
	})
}
