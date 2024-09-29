package handlers

import (
	"STUOJ/db"
	"STUOJ/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Test(c *gin.Context) {
	c.JSON(http.StatusOK, model.Response{
		Code: 1,
		Msg:  "OK",
		Data: "Hello, World!",
	})
}

func TestDb(c *gin.Context) {
	users := db.GetAllUsers()

	if users == nil {
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
