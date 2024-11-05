package user

import (
	"STUOJ/db"
	"STUOJ/model"
	"STUOJ/utils"
	"STUOJ/yuki"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateUserAvatar(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "文件上传失败", Data: err})
		return
	}
	dst := fmt.Sprintf("tmp/%s", utils.GetRandKey())
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{Code: 0, Msg: "文件上传失败", Data: err})
		return
	}
	image, err := yuki.UploadImage(dst, model.RoleAvatar)
	_ = os.Remove(dst)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code: 0,
			Msg:  "上传失败",
			Data: err,
		})
		return
	}
	id, err := utils.GetTokenUid(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code: 0,
			Msg:  "获取用户id失败",
			Data: err,
		})
		return
	}
	oldAvatarUrl, err := db.SelectUserAvatarById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code: 0,
			Msg:  "获取用户头像失败",
			Data: nil,
		})
		return
	}
	err = yuki.DeleteOldAvatar(oldAvatarUrl)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code: 0,
			Msg:  "删除旧头像失败",
			Data: err,
		})
		return
	}
	err = db.UpdateUserAvatarById(id, image.Url)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code: 0,
			Msg:  "更新用户头像失败",
			Data: err,
		})
		return
	}
	c.JSON(http.StatusOK, model.Response{
		Code: 1,
		Msg:  "更新成功",
		Data: image.Url,
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
		return
	}
	uid := uint64(id)
	avatar, err := db.SelectUserAvatarById(uid)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code: 0,
			Msg:  "获取用户头像失败",
			Data: nil,
		})
		return
	}
	c.JSON(http.StatusOK, model.Response{
		Code: 1,
		Msg:  "获取成功",
		Data: avatar,
	})
}

func ThisUserAvatar(c *gin.Context) {
	id, err := utils.GetTokenUid(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code: 0,
			Msg:  "获取用户id失败",
			Data: nil,
		})
		return
	}
	avatar, err := db.SelectUserAvatarById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code: 0,
			Msg:  "获取用户头像失败",
			Data: nil,
		})
		return
	}
	c.JSON(http.StatusOK, model.Response{
		Code: 1,
		Msg:  "获取成功",
		Data: avatar,
	})
}
