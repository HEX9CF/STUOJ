package routes

import (
	"STUOJ/server/handler"

	"github.com/gin-gonic/gin"
)

func InitProblemRoute(ginServer *gin.Engine) {
	problemPublicRoute := ginServer.Group("/problem")
	{
		problemPublicRoute.GET("/", handler.ProblemList)
		problemPublicRoute.GET("/:id", handler.ProblemInfo)

		problemPublicRoute.GET("/tag", handler.TagList)
	}
}
