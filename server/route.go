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
		c.JSON(http.StatusOK, model.RespOk("OK", "STUOJ后端启动成功！"))
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
		problemPublicRoute.GET("/difficulty/:id", handler.ProblemPublicListByDifficulty)
		problemPublicRoute.GET("/tag/:id", handler.ProblemPublicListByTagId)
		problemPublicRoute.POST("/title", handler.ProblemPublicListByTitle)
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

func InitAdminRoute() {
	adminPrivateRoute := ginServer.Group("/admin")
	{
		// 使用中间件
		adminPrivateRoute.Use(middlewares.TokenAuthAdmin())

		{
			adminPrivateRoute.GET("/user", admin.AdminUserList)
			adminPrivateRoute.GET("/user/:id", admin.AdminUserInfo)
			adminPrivateRoute.GET("/user/role/:id", admin.AdminUserListByRole)
			adminPrivateRoute.POST("/user", admin.AdminUserAdd)
			adminPrivateRoute.PUT("/user", admin.AdminUserModify)
			adminPrivateRoute.DELETE("/user/:id", admin.AdminUserRemove)
		}
		{
			adminPrivateRoute.GET("/problem", admin.AdminProblemList)
			adminPrivateRoute.GET("/problem/status/:id", admin.AdminProblemListByStatus)
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
			adminPrivateRoute.GET("/statistics", admin.AdminStatisticsList)
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
