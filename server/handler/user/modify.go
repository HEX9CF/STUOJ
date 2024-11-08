package user

import (
	"STUOJ/internal/db"
	"STUOJ/internal/model"
	"STUOJ/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 修改用户信息
type ReqUserModify struct {
	Username  string `json:"username" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Signature string `json:"signature" binding:"required"`
}

func UserModify(c *gin.Context) {
	var req ReqUserModify

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "参数错误",
			Data: nil,
		})
		return
	}

	// 获取用户id
	id, err := utils.GetTokenUid(c)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "用户未登录",
			Data: nil,
		})
		return
	}

	// 读取用户
	u, err := db.SelectUserById(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "修改失败，用户不存在",
			Data: nil,
		})
		return
	}

	// 修改用户
	u.Username = req.Username
	u.Email = req.Email
	u.Signature = req.Signature

	err = db.UpdateUserByIdExceptPassword(u)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "修改失败, 用户名或邮箱已存在",
			Data: nil,
		})
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.Response{
		Code: model.ResponseCodeOk,
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
			Code: model.ResponseCodeError,
			Msg:  "参数错误",
			Data: nil,
		})
		return
	}

	// 获取用户id
	id, err := utils.GetTokenUid(c)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "用户未登录",
			Data: nil,
		})
		return
	}

	// 读取用户
	u, err := db.SelectUserById(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "修改失败，用户不存在",
			Data: nil,
		})
		return
	}

	// 修改用户密码
	u.Password = req.Password

	err = db.UpdateUserPasswordById(u)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "修改失败",
			Data: nil,
		})
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model.Response{
		Code: model.ResponseCodeOk,
		Msg:  "修改成功",
		Data: nil,
	})
}
