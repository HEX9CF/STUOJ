package route

import (
	handler_admin "STUOJ/server/handler-admin"
	"STUOJ/server/middlewares"
	"github.com/gin-gonic/gin"
)

func InitAdminRoute(ginServer *gin.Engine) {
	adminPrivateRoute := ginServer.Group("/admin")
	{
		// 使用中间件
		adminPrivateRoute.Use(middlewares.TokenAuthAdmin())

		{
			adminPrivateRoute.GET("/user", handler_admin.AdminUserList)
			adminPrivateRoute.GET("/user/:id", handler_admin.AdminUserInfo)
			adminPrivateRoute.GET("/user/role/:id", handler_admin.AdminUserListOfRole)
			adminPrivateRoute.POST("/user", handler_admin.AdminUserAdd)
			adminPrivateRoute.PUT("/user", handler_admin.AdminUserModify)
			adminPrivateRoute.DELETE("/user/:id", handler_admin.AdminUserRemove)
		}
		{
			adminPrivateRoute.GET("/problem", handler_admin.AdminProblemList)
			adminPrivateRoute.GET("/problem/status/:id", handler_admin.AdminProblemListOfStatus)
			adminPrivateRoute.GET("/problem/:id", handler_admin.AdminProblemInfo)
			adminPrivateRoute.POST("/problem", handler_admin.AdminProblemAdd)
			adminPrivateRoute.PUT("/problem", handler_admin.AdminProblemModify)
			adminPrivateRoute.DELETE("/problem/:id", handler_admin.AdminProblemRemove)
			adminPrivateRoute.POST("/problem/tag", handler_admin.AdminProblemAddTag)
			adminPrivateRoute.DELETE("/problem/tag", handler_admin.AdminProblemRemoveTag)
			adminPrivateRoute.POST("/problem/fps", handler_admin.AdminProblemParseFromFps)

		}
		{
			adminPrivateRoute.GET("/testcase/:id", handler_admin.AdminTestcaseInfo)
			adminPrivateRoute.POST("/testcase", handler_admin.AdminTestcaseAdd)
			adminPrivateRoute.PUT("/testcase", handler_admin.AdminTestcaseModify)
			adminPrivateRoute.DELETE("/testcase/:id", handler_admin.AdminTestcaseRemove)
			adminPrivateRoute.POST("/testcase/datamake", handler_admin.AdminTestcaseDataMake)
		}
		{
			adminPrivateRoute.GET("/tag", handler_admin.AdminTagList)
			adminPrivateRoute.POST("/tag", handler_admin.AdminTagAdd)
			adminPrivateRoute.PUT("/tag", handler_admin.AdminTagModify)
			adminPrivateRoute.DELETE("/tag/:id", handler_admin.AdminTagRemove)
		}
		{
			adminPrivateRoute.GET("/solution/:id", handler_admin.AdminSolutionInfo)
			adminPrivateRoute.POST("/solution", handler_admin.AdminSolutionAdd)
			adminPrivateRoute.PUT("/solution", handler_admin.AdminSolutionModify)
			adminPrivateRoute.DELETE("/solution/:id", handler_admin.AdminSolutionRemove)
		}
		{
			adminPrivateRoute.GET("/record", handler_admin.AdminRecordList)
			adminPrivateRoute.GET("/record/:id", handler_admin.AdminRecordInfo)
			adminPrivateRoute.DELETE("/record/:id", handler_admin.AdminRecordRemove)
		}
		{
			adminPrivateRoute.GET("/blog", handler_admin.AdminBlogList)
			adminPrivateRoute.GET("/blog/status/:id", handler_admin.AdminBlogListOfStatus)
			adminPrivateRoute.GET("/blog/:id", handler_admin.AdminBlogInfo)
			adminPrivateRoute.POST("/blog", handler_admin.AdminBlogAdd)
			adminPrivateRoute.PUT("/blog", handler_admin.AdminBlogModify)
			adminPrivateRoute.DELETE("/blog/:id", handler_admin.AdminBlogRemove)
		}
		{
			adminPrivateRoute.POST("/comment", handler_admin.AdminCommentAdd)
			adminPrivateRoute.PUT("/comment", handler_admin.AdminCommentModify)
			adminPrivateRoute.DELETE("/comment/:id", handler_admin.AdminCommentRemove)
		}
		{
			adminPrivateRoute.GET("/statistics/user", handler_admin.AdminStatisticsUser)
			adminPrivateRoute.GET("/statistics/user/role", handler_admin.AdminStatisticsUserOfRole)
			adminPrivateRoute.GET("/statistics/user/register", handler_admin.AdminStatisticsUserOfRegister)

			adminPrivateRoute.GET("/statistics/tag", handler_admin.AdminStatisticsTag)
			adminPrivateRoute.GET("/statistics/problem", handler_admin.AdminStatisticsProblem)
			adminPrivateRoute.GET("/statistics/problem/insert", handler_admin.AdminStatisticsProblemOfInsert)
			adminPrivateRoute.GET("/statistics/problem/update", handler_admin.AdminStatisticsProblemOfUpdate)
			adminPrivateRoute.GET("/statistics/problem/delete", handler_admin.AdminStatisticsProblemOfDelete)

			adminPrivateRoute.GET("/statistics/judge", handler_admin.AdminStatisticsJudge)

			adminPrivateRoute.GET("/statistics/record", handler_admin.AdminStatisticsRecord)
			adminPrivateRoute.GET("/statistics/record/submit", handler_admin.AdminStatisticsRecordOfSubmit)
			adminPrivateRoute.GET("/statistics/record/language", handler_admin.AdminStatisticsRecordOfLanguage)
			adminPrivateRoute.GET("/statistics/submission/status", handler_admin.AdminStatisticsSubmissionOfStatus)
			adminPrivateRoute.GET("/statistics/judgement/status", handler_admin.AdminStatisticsJudgementOfStatus)

			adminPrivateRoute.GET("/statistics/blog", handler_admin.AdminStatisticsBlog)
			adminPrivateRoute.GET("/statistics/blog/submit", handler_admin.AdminStatisticsBlogOfSubmit)
			adminPrivateRoute.GET("/statistics/comment/submit", handler_admin.AdminStatisticsCommentOfSubmit)
		}
	}

	rootPrivateRoute := ginServer.Group("/admin")
	{
		// 使用中间件
		rootPrivateRoute.Use(middlewares.TokenAuthRoot())

		rootPrivateRoute.PUT("/user/role", handler_admin.AdminUserModifyRole)

		rootPrivateRoute.GET("/config", handler_admin.AdminConfigList)
		//rootPrivateRoute.PUT("/config", handler-admin.AdminConfigModify)
	}
}
