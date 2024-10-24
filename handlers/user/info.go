package user

import (
	"STUOJ/db"
	"STUOJ/model"
	"STUOJ/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

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
