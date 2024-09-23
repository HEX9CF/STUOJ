package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	ginServer := gin.Default()

	ginServer.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "STUOJ后端服务启动成功！",
		})
	})

	err := ginServer.Run()
	if err != nil {
		return
	}
}
