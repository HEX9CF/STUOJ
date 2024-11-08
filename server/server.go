package server

import (
	"STUOJ/server/handler"
	"github.com/gin-gonic/gin"
)

var (
	ginServer *gin.Engine
)

func InitServer() error {
	ginServer = gin.Default()
	err := handler.InitRoute()
	if err != nil {
		return err
	}

	return nil
}
