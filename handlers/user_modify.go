package handlers

import (
	"STUOJ/db"
	"STUOJ/model"
	"STUOJ/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

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
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, model.Response{
			Code: 0,
			Msg:  "用户未登录",
			Data: nil,
		})
		return
	}

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
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, model.Response{
			Code: 0,
			Msg:  "用户未登录",
			Data: nil,
		})
		return
	}

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
