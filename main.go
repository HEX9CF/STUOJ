package main

import (
	"github.com/gin-gonic/gin"
)

const (
	PORT = ":8080"
)

func main() {
	ginServer := gin.Default()

	ginServer.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "STUOJ后端服务启动成功！",
		})
	})

	err := ginServer.Run(PORT)
	if err != nil {
		return
	}
}
