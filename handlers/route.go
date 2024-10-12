package handlers

import (
	"STUOJ/handlers/user"
	"STUOJ/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRoute() {
	// index
	ginServer.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, model.Response{
			Code: 1,
			Msg:  "OK",
			Data: "STUOJ后端启动成功！",
		})
	})

	// 404
	ginServer.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, model.Response{
			Code: 0,
			Msg:  "404 Not Found",
			Data: nil,
		})
	})

	// 初始化路由
	InitUserRoute()
	InitTestRoute()

	// 启动服务
	err := ginServer.Run(PORT)
	if err != nil {
		return
	}
}

func InitUserRoute() {
	userRoute := ginServer.Group("/user")
	{
		userRoute.POST("/login", user.UserLogin)
		userRoute.POST("/register", user.UserRegister)
		userRoute.POST("/logout", user.UserLogout)
		userRoute.POST("/data", user.UserData)
	}
}

func InitTestRoute() {
	testRoute := ginServer.Group("/test")
	{
		testRoute.GET("/", Test)
		testRoute.GET("/db", TestDb)
	}
}
