package handler

import (
	"STUOJ/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	c.JSON(http.StatusOK, model.RespOk("OK", "Hello, World!"))
}
