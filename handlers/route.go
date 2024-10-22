package handlers

import (
	"STUOJ/conf"
	"STUOJ/middlewares"
	"STUOJ/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRoute() error {
	config := conf.Conf.Server

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
	InitTestRoute()
	InitUserRoute()
	InitProblemRoute()
	InitJudgeRoute()
	InitRecordRoute()

	// 启动服务
	err := ginServer.Run(":" + config.Port)
	if err != nil {
		return err
	}

	return nil
}

func InitTestRoute() {
	testRoute := ginServer.Group("/test")
	{
		testRoute.GET("/", Test)
	}
}

func InitUserRoute() {
	userPublicRoute := ginServer.Group("/user")
	{
		userPublicRoute.GET("/", UserList)
		userPublicRoute.GET("/:id", UserInfo)
		userPublicRoute.POST("/login", UserLogin)
		userPublicRoute.POST("/register", UserRegister)
	}
	userProtectedRoute := ginServer.Group("/user")
	{
		userProtectedRoute.Use(middlewares.TokenAuth())
		userProtectedRoute.GET("/current", UserCurrentId)
		userProtectedRoute.PUT("/modify", UserModify)
		userProtectedRoute.PUT("/password", UserChangePassword)
	}
}

func InitProblemRoute() {
	problemPublicRoute := ginServer.Group("/problem")
	{
		problemPublicRoute.GET("/", ProblemList)
		problemPublicRoute.GET("/:id", ProblemInfo)
	}
}

func InitJudgeRoute() {
	judgePublicRoute := ginServer.Group("/judge")
	{
		judgePublicRoute.GET("/language", JudgeLanguageList)
	}
}

func InitRecordRoute() {
	recordPublicRoute := ginServer.Group("/record")
	{
		recordPublicRoute.GET("/", RecordList)
		recordPublicRoute.GET("/:id", RecordInfo)
		recordPublicRoute.GET("/user/:id", RecordListOfUser)
		recordPublicRoute.GET("/problem/:id", RecordListOfProblem)
	}
}
