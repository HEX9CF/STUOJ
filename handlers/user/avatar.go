package user

import (
	"STUOJ/db"
	"STUOJ/lskypro"
	"STUOJ/model"
	"STUOJ/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

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
