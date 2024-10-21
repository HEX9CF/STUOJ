package handlers

import (
	"github.com/gin-gonic/gin"
)

var (
	ginServer *gin.Engine
)

func Init() error {
	ginServer = gin.Default()
	err := InitRoute()
	if err != nil {
		return err
	}

	return nil
}
