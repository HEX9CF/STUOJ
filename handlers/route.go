package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"

)



func InitRoute() {
	// index
	ginServer.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "STUOJ后端服务启动成功！",
		})
	})

	// 404
	ginServer.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "404 Not Found",
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
	userRoute:=ginServer.Group("/user")
	{
		userRoute.POST("/login",UserLogin)
		userRoute.POST("/register",UserRegister)
		userRoute.POST("/logout",UserLogout)
		userRoute.POST("/data",UserData)
	}

}
