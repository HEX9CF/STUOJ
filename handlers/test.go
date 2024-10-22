package handlers

import (
	"STUOJ/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	c.JSON(http.StatusOK, model.Response{
		Code: 1,
		Msg:  "OK",
		Data: "Hello, World!",
	})
}
