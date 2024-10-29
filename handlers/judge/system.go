package judge

import (
	"STUOJ/database"
	"STUOJ/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 获取语言列表
func JudgeLanguageList(c *gin.Context) {
	languages, err := database.SelectAllLanguages()
	if err != nil || languages == nil {
		if err != nil {
			log.Println(err)
		}
		c.JSON(http.StatusOK, model.Response{
			Code: model.ResponseCodeError,
			Msg:  "获取失败",
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code: model.ResponseCodeOk,
		Msg:  "OK",
		Data: languages,
	})
}
