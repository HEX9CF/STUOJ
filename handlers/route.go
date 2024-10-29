package handlers

import (
	"STUOJ/conf"
	"STUOJ/handlers/admin"
	"STUOJ/handlers/judge"
	"STUOJ/handlers/user"
	"STUOJ/middlewares"
	"STUOJ/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoute() error {
	config := conf.Conf.Server

	// index
	ginServer.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, model.Response{
			Code: model.ResponseCodeOk,
			Msg:  "OK",
			Data: "STUOJ后端启动成功！",
		})
	})

	// 404
	ginServer.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, model.Response{
			Code: model.ResponseCodeError,
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
		testRoute.GET("/", Test)
	}
}

func InitUserRoute() {
	userPublicRoute := ginServer.Group("/user")
	{
		userPublicRoute.GET("/", user.UserList)
		userPublicRoute.GET("/avatar/:id", user.UserAvatar)
		userPublicRoute.GET("/:id", user.UserInfo)
		userPublicRoute.POST("/login", user.UserLogin)
		userPublicRoute.POST("/register", user.UserRegister)
	}
	userProtectedRoute := ginServer.Group("/user")
	{
		// 使用中间件
		userProtectedRoute.Use(middlewares.TokenAuthUser())

		userProtectedRoute.GET("/current", user.UserCurrentId)
		userPublicRoute.GET("/avatar", user.ThisUserAvatar)
		userProtectedRoute.PUT("/modify", user.UserModify)
		userProtectedRoute.PUT("/password", user.UserChangePassword)
		userProtectedRoute.POST("/avatar", user.UpdateUserAvatar)
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
		judgePublicRoute.GET("/language", judge.JudgeLanguageList)
	}
	judgePrivateRoute := ginServer.Group("/judge")
	{
		// 使用中间件
		judgePrivateRoute.Use(middlewares.TokenAuthUser())

		judgePrivateRoute.POST("/submit", judge.JudgeSubmit)
	}
}

func InitRecordRoute() {
	recordPublicRoute := ginServer.Group("/record")
	{
		recordPublicRoute.GET("/", RecordList)
		recordPublicRoute.GET("/:id", RecordInfo)
		recordPublicRoute.GET("/user/:id", RecordListOfUser)
		recordPublicRoute.GET("/problem/:id", RecordListOfProblem)
		recordPublicRoute.GET("/point/problem/:id", RecordPointListOfProblem)
	}
}

func InitAdminRoute() {
	adminPrivateRoute := ginServer.Group("/admin")
	{
		// 使用中间件
		adminPrivateRoute.Use(middlewares.TokenAuthAdmin())

		adminPrivateRoute.GET("/user", admin.AdminUserList)
		adminPrivateRoute.GET("/user/:id", admin.AdminUserInfo)
		adminPrivateRoute.POST("/user", admin.AdminUserAdd)
		adminPrivateRoute.PUT("/user", admin.AdminUserModify)
		adminPrivateRoute.DELETE("/user/:id", admin.AdminUserRemove)

		adminPrivateRoute.GET("/problem", admin.AdminProblemList)
		adminPrivateRoute.GET("/problem/:id", admin.AdminProblemInfo)
		adminPrivateRoute.POST("/problem", admin.AdminProblemAdd)
		adminPrivateRoute.PUT("/problem", admin.AdminProblemModify)
		adminPrivateRoute.DELETE("/problem/:id", admin.AdminProblemRemove)

		//adminPrivateRoute.GET("/record", AdminRecord)

		//adminPrivateRoute.GET("/point", AdminPoint)

		//adminPrivateRoute.GET("/system", AdminSystem)
	}
}
