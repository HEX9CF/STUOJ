package handlers

import (
	"STUOJ/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "OK",
		"data": "Hello, world!",
	})
}

func TestDb(c *gin.Context) {
	users := db.GetAllUsers()
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "OK",
		"data": users,
	})
}
