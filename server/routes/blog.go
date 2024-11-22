package routes

import (
	"STUOJ/server/handler"
	"STUOJ/server/middlewares"
	"github.com/gin-gonic/gin"
)

func InitBlogRoute(ginServer *gin.Engine) {
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
