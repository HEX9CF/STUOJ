package handlers

import (
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
