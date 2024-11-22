package routes

import (
	"STUOJ/server/handler"
	"github.com/gin-gonic/gin"
)

func InitProblemRoute(ginServer *gin.Engine) {
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
