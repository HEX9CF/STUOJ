package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRoute() {
	// index
	ginServer.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "STUOJ后端服务启动成功！",
			"data": nil,
		})
	})

	// 404
	ginServer.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 0,
			"msg":  "404 Not Found",
			"data": nil,
		})
	})
	InitUserRoute()
	// 启动服务
	err := ginServer.Run(PORT)
	if err != nil {
		return
	}
}

func InitUserRoute() {
	userRoute := ginServer.Group("/user")
	{
		userRoute.POST("/login", UserLogin)
		userRoute.POST("/register", UserRegister)
		userRoute.POST("/logout", UserLogout)
		userRoute.POST("/data", UserData)
	}
}
