package routes

import (
	"STUOJ/server/handler"
	"STUOJ/server/middlewares"
	"github.com/gin-gonic/gin"
)

func InitJudgeRoute(ginServer *gin.Engine) {
	judgePublicRoute := ginServer.Group("/judge")
	{
		judgePublicRoute.GET("/language", handler.JudgeLanguageList)
	}
	judgePrivateRoute := ginServer.Group("/judge")
	{
		// 使用中间件
		judgePrivateRoute.Use(middlewares.TokenAuthUser())

		judgePrivateRoute.POST("/submit", handler.JudgeSubmit)
		judgePrivateRoute.POST("/testrun", handler.JudgeTestRun)
	}
}
