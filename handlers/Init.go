package handlers

import(
	"github.com/gin-gonic/gin"
)

const (
	PORT = ":8080" // 端口
)

var (
	ginServer *gin.Engine
)

func Init(){
	ginServer = gin.Default()
	InitRoute()
}