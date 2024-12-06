package routes

import (
	"STUOJ/server/handler"
	"STUOJ/server/middlewares"

	"github.com/gin-gonic/gin"
)

func InitBlogRoute(ginServer *gin.Engine) {
	blogPublicRoute := ginServer.Group("/blog")
	{
		blogPublicRoute.GET("/", handler.BlogList)
		blogPublicRoute.GET("/:id", handler.BlogInfo)
	}
	blogPrivateRoute := ginServer.Group("/blog")
	{
		// 使用中间件
		blogPrivateRoute.Use(middlewares.TokenAuthUser())

		blogPrivateRoute.POST("/", handler.BlogUpload)
		blogPrivateRoute.PUT("/", handler.BlogEdit)
		blogPrivateRoute.PUT("/:id", handler.BlogSubmit)
		blogPrivateRoute.DELETE("/:id", handler.BlogRemove)
	}
}
