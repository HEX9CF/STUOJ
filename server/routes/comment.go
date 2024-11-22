package routes

import (
	"STUOJ/server/handler"
	"STUOJ/server/middlewares"
	"github.com/gin-gonic/gin"
)

func InitCommentRoute(ginServer *gin.Engine) {
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
