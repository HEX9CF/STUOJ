package admin

import (
	"STUOJ/internal/db"
	model2 "STUOJ/internal/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// 获取用户信息
func AdminUserInfo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model2.Response{
			Code: model2.ResponseCodeError,
			Msg:  "参数错误",
			Data: nil,
		})
		return
	}

	uid := uint64(id)
	user, err := db.SelectUserById(uid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model2.Response{
			Code: model2.ResponseCodeError,
			Msg:  "获取用户信息失败",
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, model2.Response{
		Code: model2.ResponseCodeOk,
		Msg:  "OK",
		Data: user,
	})
}

// 获取用户列表
func AdminUserList(c *gin.Context) {
	users, err := db.SelectAllUsers()
	if err != nil || users == nil {
		if err != nil {
			log.Println(err)
		}
		c.JSON(http.StatusOK, model2.Response{
			Code: model2.ResponseCodeError,
			Msg:  "获取失败",
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, model2.Response{
		Code: model2.ResponseCodeOk,
		Msg:  "OK",
		Data: users,
	})
}

// 根据角色获取用户列表
func AdminUserListByRole(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model2.Response{
			Code: model2.ResponseCodeError,
			Msg:  "参数错误",
			Data: nil,
		})
		return
	}

	rid := model2.UserRole(id)
	users, err := db.SelectUsersByRole(rid)
	if err != nil || users == nil {
		if err != nil {
			log.Println(err)
		}
		c.JSON(http.StatusOK, model2.Response{
			Code: model2.ResponseCodeError,
			Msg:  "获取失败",
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, model2.Response{
		Code: model2.ResponseCodeOk,
		Msg:  "OK",
		Data: users,
	})
}

// 添加普通用户
type ReqUserAdd struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Avatar    string `json:"avatar" binding:"required"`
	Signature string `json:"signature" binding:"required"`
}

func AdminUserAdd(c *gin.Context) {
	var req ReqUserAdd

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model2.Response{
			Code: model2.ResponseCodeError,
			Msg:  "参数错误",
			Data: nil,
		})
		return
	}

	// 初始化用户
	u := model2.User{
		Username:  req.Username,
		Password:  req.Password,
		Email:     req.Email,
		Avatar:    req.Avatar,
		Signature: req.Signature,
	}
	u.Id, err = db.InsertUser(u)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model2.Response{
			Code: model2.ResponseCodeError,
			Msg:  "添加失败，用户名或邮箱已存在",
			Data: nil,
		})
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model2.Response{
		Code: model2.ResponseCodeOk,
		Msg:  "添加成功，返回用户ID",
		Data: u.Id,
	})
}

// 修改用户
type ReqUserModify struct {
	Id        uint64 `json:"id" binding:"required"`
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Avatar    string `json:"avatar" binding:"required"`
	Signature string `json:"signature" binding:"required"`
}

func AdminUserModify(c *gin.Context) {
	var req ReqUserModify

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model2.Response{
			Code: model2.ResponseCodeError,
			Msg:  "参数错误",
			Data: nil,
		})
		return
	}

	// 读取用户
	u, err := db.SelectUserById(req.Id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model2.Response{
			Code: model2.ResponseCodeError,
			Msg:  "修改失败，用户不存在",
			Data: nil,
		})
		return
	}

	// 修改用户
	u.Username = req.Username
	u.Password = req.Password
	u.Email = req.Email
	u.Avatar = req.Avatar
	u.Signature = req.Signature

	err = db.UpdateUserById(u)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model2.Response{
			Code: model2.ResponseCodeError,
			Msg:  "修改失败",
			Data: nil,
		})
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model2.Response{
		Code: model2.ResponseCodeOk,
		Msg:  "修改成功",
		Data: nil,
	})
}

// 删除用户
func AdminUserRemove(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model2.Response{
			Code: model2.ResponseCodeError,
			Msg:  "参数错误",
			Data: nil,
		})
		return
	}

	uid := uint64(id)
	_, err = db.SelectUserById(uid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model2.Response{
			Code: model2.ResponseCodeError,
			Msg:  "删除失败，用户不存在",
			Data: nil,
		})
		return
	}

	err = db.DeleteUserById(uid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model2.Response{
			Code: model2.ResponseCodeError,
			Msg:  "删除失败",
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, model2.Response{
		Code: model2.ResponseCodeOk,
		Msg:  "删除成功",
		Data: nil,
	})
}

// 设置用户角色
type ReqUserModifyRole struct {
	Id   uint64          `json:"id" binding:"required"`
	Role model2.UserRole `json:"role" binding:"required"`
}

func AdminUserModifyRole(c *gin.Context) {
	var req ReqUserModifyRole

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model2.Response{
			Code: model2.ResponseCodeError,
			Msg:  "参数错误",
			Data: nil,
		})
		return
	}

	// 读取用户
	u, err := db.SelectUserById(req.Id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model2.Response{
			Code: model2.ResponseCodeError,
			Msg:  "修改失败，用户不存在",
			Data: nil,
		})
		return
	}

	// 修改用户
	u.Role = req.Role

	err = db.UpdateUserRoleById(u)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model2.Response{
			Code: model2.ResponseCodeError,
			Msg:  "修改失败",
			Data: nil,
		})
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, model2.Response{
		Code: model2.ResponseCodeOk,
		Msg:  "修改成功",
		Data: nil,
	})
}
