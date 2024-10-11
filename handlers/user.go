package handlers

import (
	"STUOJ/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type ReqUserLogin struct {
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func UserLogin(c *gin.Context) {

}

type ReqUserRegister struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func UserRegister(c *gin.Context) {
	var req ReqUserRegister

	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.Response{
			Code: 0,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code: 1,
		Msg:  "OK",
		Data: req,
	})
}

func UserLogout(c *gin.Context) {
}

func UserData(c *gin.Context) {
}
