package server

import (
	"STUOJ/internal/conf"
	"STUOJ/internal/model"
	"STUOJ/server/handler"
	"STUOJ/server/handler/admin"
	"STUOJ/server/middlewares"
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
	InitAdminRoute()

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

func InitAdminRoute() {
	adminPrivateRoute := ginServer.Group("/admin")
	{
		// 使用中间件
		adminPrivateRoute.Use(middlewares.TokenAuthAdmin())

		{
			adminPrivateRoute.GET("/user", admin.AdminUserList)
			adminPrivateRoute.GET("/user/:id", admin.AdminUserInfo)
			adminPrivateRoute.GET("/user/role/:id", admin.AdminUserListOfRole)
			adminPrivateRoute.POST("/user", admin.AdminUserAdd)
			adminPrivateRoute.PUT("/user", admin.AdminUserModify)
			adminPrivateRoute.DELETE("/user/:id", admin.AdminUserRemove)
		}
		{
			adminPrivateRoute.GET("/problem", admin.AdminProblemList)
			adminPrivateRoute.GET("/problem/status/:id", admin.AdminProblemListOfStatus)
			adminPrivateRoute.GET("/problem/:id", admin.AdminProblemInfo)
			adminPrivateRoute.POST("/problem", admin.AdminProblemAdd)
			adminPrivateRoute.PUT("/problem", admin.AdminProblemModify)
			adminPrivateRoute.DELETE("/problem/:id", admin.AdminProblemRemove)
			adminPrivateRoute.POST("/problem/tag", admin.AdminProblemAddTag)
			adminPrivateRoute.DELETE("/problem/tag", admin.AdminProblemRemoveTag)
			adminPrivateRoute.POST("/problem/fps", admin.AdminProblemParseFromFps)

		}
		{
			adminPrivateRoute.GET("/testcase/:id", admin.AdminTestcaseInfo)
			adminPrivateRoute.POST("/testcase", admin.AdminTestcaseAdd)
			adminPrivateRoute.PUT("/testcase", admin.AdminTestcaseModify)
			adminPrivateRoute.DELETE("/testcase/:id", admin.AdminTestcaseRemove)
			adminPrivateRoute.POST("/testcase/datamake", admin.AdminTestcaseDataMake)
		}
		{
			adminPrivateRoute.GET("/tag", admin.AdminTagList)
			adminPrivateRoute.POST("/tag", admin.AdminTagAdd)
			adminPrivateRoute.PUT("/tag", admin.AdminTagModify)
			adminPrivateRoute.DELETE("/tag/:id", admin.AdminTagRemove)
		}
		{
			adminPrivateRoute.GET("/solution/:id", admin.AdminSolutionInfo)
			adminPrivateRoute.POST("/solution", admin.AdminSolutionAdd)
			adminPrivateRoute.PUT("/solution", admin.AdminSolutionModify)
			adminPrivateRoute.DELETE("/solution/:id", admin.AdminSolutionRemove)
		}
		{
			adminPrivateRoute.GET("/record", admin.AdminRecordList)
			adminPrivateRoute.GET("/record/:id", admin.AdminRecordInfo)
			adminPrivateRoute.DELETE("/record/:id", admin.AdminRecordRemove)
		}
		{
			adminPrivateRoute.GET("/blog", admin.AdminBlogList)
			adminPrivateRoute.GET("/blog/status/:id", admin.AdminBlogListOfStatus)
			adminPrivateRoute.GET("/blog/:id", admin.AdminBlogInfo)
			adminPrivateRoute.POST("/blog", admin.AdminBlogAdd)
			adminPrivateRoute.PUT("/blog", admin.AdminBlogModify)
			adminPrivateRoute.DELETE("/blog/:id", admin.AdminBlogRemove)
		}
		{
			adminPrivateRoute.POST("/comment", admin.AdminCommentAdd)
			adminPrivateRoute.PUT("/comment", admin.AdminCommentModify)
			adminPrivateRoute.DELETE("/comment/:id", admin.AdminCommentRemove)
		}
		{
			adminPrivateRoute.GET("/statistics/tag", admin.AdminStatisticsTag)
			adminPrivateRoute.GET("/statistics/judge", admin.AdminStatisticsJudge)
			adminPrivateRoute.GET("/statistics/user", admin.AdminStatisticsUser)
			adminPrivateRoute.GET("/statistics/user/role", admin.AdminStatisticsUserOfRole)
			adminPrivateRoute.GET("/statistics/user/register", admin.AdminStatisticsUserOfRegister)
			adminPrivateRoute.GET("/statistics/problem", admin.AdminStatisticsProblem)
			adminPrivateRoute.GET("/statistics/record", admin.AdminStatisticsRecord)
			adminPrivateRoute.GET("/statistics/blog", admin.AdminStatisticsBlog)
			adminPrivateRoute.GET("/statistics/comment", admin.AdminStatisticsComment)
		}
	}

	rootPrivateRoute := ginServer.Group("/admin")
	{
		// 使用中间件
		rootPrivateRoute.Use(middlewares.TokenAuthRoot())

		rootPrivateRoute.PUT("/user/role", admin.AdminUserModifyRole)

		rootPrivateRoute.GET("/config", admin.AdminConfigList)
		//rootPrivateRoute.PUT("/config", admin.AdminConfigModify)
	}
}
