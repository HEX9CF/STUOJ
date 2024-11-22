package server

import (
	"STUOJ/internal/conf"
	"STUOJ/internal/model"
	"STUOJ/server/handler"
	"STUOJ/server/middlewares"
	"STUOJ/server/route"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoute() error {
	config := conf.Conf.Server

	// index
	ginServer.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, model.RespOk("STUOJ后端启动成功！", nil))
	})

	// 404
	ginServer.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, model.RespError("404 Not Found", nil))
	})

	// 初始化路由
	InitTestRoute()
	InitUserRoute()
	InitProblemRoute()
	InitJudgeRoute()
	InitRecordRoute()
	InitBlogRoute()
	InitCommentRoute()
	route.InitAdminRoute(ginServer)

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
		testRoute.GET("/", handler.Test)
	}
}

func InitUserRoute() {
	userPublicRoute := ginServer.Group("/user")
	{
		userPublicRoute.GET("/:id", handler.UserInfo)
		userPublicRoute.POST("/login", handler.UserLogin)
		userPublicRoute.POST("/register", handler.UserRegister)
	}
	userProtectedRoute := ginServer.Group("/user")
	{
		// 使用中间件
		userProtectedRoute.Use(middlewares.TokenAuthUser())

		userProtectedRoute.GET("/current", handler.UserCurrentId)
		userProtectedRoute.PUT("/modify", handler.UserModify)
		userProtectedRoute.PUT("/password", handler.UserChangePassword)
		userProtectedRoute.POST("/avatar", handler.ModifyUserAvatar)
	}
}

func InitProblemRoute() {
	problemPublicRoute := ginServer.Group("/problem")
	{
		problemPublicRoute.GET("/", handler.ProblemPublicList)
		problemPublicRoute.GET("/difficulty/:id", handler.ProblemPublicListOfDifficulty)
		problemPublicRoute.GET("/tag/:id", handler.ProblemPublicListOfTagId)
		problemPublicRoute.POST("/title", handler.ProblemPublicListOfTitle)
		problemPublicRoute.GET("/:id", handler.ProblemPublicInfo)

		problemPublicRoute.GET("/tag", handler.TagList)
	}
}

func InitJudgeRoute() {
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

func InitRecordRoute() {
	recordPublicRoute := ginServer.Group("/record")
	{
		recordPublicRoute.GET("/", handler.RecordList)
		recordPublicRoute.GET("/:id", handler.RecordInfo)
		recordPublicRoute.GET("/user/:id", handler.RecordListOfUser)
		recordPublicRoute.GET("/problem/:id", handler.RecordListOfProblem)
	}
}

func InitBlogRoute() {
	blogPublicRoute := ginServer.Group("/blog")
	{
		blogPublicRoute.GET("/", handler.BlogPublicList)
		blogPublicRoute.GET("/:id", handler.BlogPublicInfo)
		blogPublicRoute.GET("/user/:id", handler.BlogPublicListOfUser)
		blogPublicRoute.GET("/draft", handler.BlogDraftListOfUser)
		blogPublicRoute.GET("/problem/:id", handler.BlogPublicListOfProblem)
		blogPublicRoute.POST("/title", handler.BlogPublicListOfTitle)
	}
	blogPrivateRoute := ginServer.Group("/blog")
	{
		// 使用中间件
		blogPrivateRoute.Use(middlewares.TokenAuthUser())

		blogPrivateRoute.POST("/", handler.BlogSave)
		blogPrivateRoute.PUT("/", handler.BlogEdit)
		blogPrivateRoute.PUT("/:id", handler.BlogSubmit)
		blogPrivateRoute.DELETE("/:id", handler.BlogRemove)
	}
}

func InitCommentRoute() {
	commentPublicRoute := ginServer.Group("/comment")
	{
		commentPublicRoute.GET("/user/:id", handler.CommentPublicListOfUser)
	}
	commentPrivateRoute := ginServer.Group("/comment")
	{
		// 使用中间件
		commentPrivateRoute.Use(middlewares.TokenAuthUser())

		commentPrivateRoute.POST("/", handler.CommentAdd)
		commentPrivateRoute.DELETE("/:id", handler.CommentRemove)
	}
}
