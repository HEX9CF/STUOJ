package server

import (
	"STUOJ/internal/model"
	"STUOJ/server/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoute() error {
	// index
	ginServer.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, model.RespOk("STUOJ后端启动成功！", nil))
	})

	// 404
	ginServer.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, model.RespError("404 Not Found", nil))
	})

	// 初始化路由
	routes.InitUserRoute(ginServer)
	routes.InitProblemRoute(ginServer)
	routes.InitJudgeRoute(ginServer)
	routes.InitRecordRoute(ginServer)
	routes.InitBlogRoute(ginServer)
	routes.InitCommentRoute(ginServer)
	routes.InitAdminRoute(ginServer)

	return nil
}
