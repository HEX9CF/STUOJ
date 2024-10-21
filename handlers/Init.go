package handlers

import (
	"github.com/gin-gonic/gin"
)

var (
	ginServer *gin.Engine
)

func Init() {
	ginServer = gin.Default()
	InitRoute()
}
