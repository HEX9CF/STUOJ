package routes

import (
	"STUOJ/server/handler"
	"STUOJ/server/middlewares"

	"github.com/gin-gonic/gin"
)

func InitMiscRoute(ginServer *gin.Engine) {
	miscRoute := ginServer.Group("/misc")
	{
		miscRoute.Use(middlewares.TokenAuthUser())
		miscRoute.POST("/uploadimage", handler.UploadImage)
	}
}
