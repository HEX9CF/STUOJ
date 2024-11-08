package middlewares

import (
	"STUOJ/internal/conf"
	model2 "STUOJ/internal/model"
	"STUOJ/internal/service/user"
	"STUOJ/server/model"
	"STUOJ/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func TokenAuthUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 验证token
		err := tokenVerify(c)
		if err != nil {
			log.Println(err)
			return
		}

		// 校验用户角色
		role, err := getUserRole(c)
		if err != nil {
			log.Println(err)
			return
		}
		switch role {
		case model2.UserRoleBanned:
			c.JSON(http.StatusUnauthorized, model.RespError("拒绝访问，用户已被封禁", nil))
			c.Abort()
			return
		default:
			break
		}

		// 自动刷新token
		err = tokenAutoRefresh(c)
		if err != nil {
			log.Println(err)
			return
		}

		// 放行
		c.Next()
	}
}

func TokenAuthAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 验证token
		err := tokenVerify(c)
		if err != nil {
			log.Println(err)
			return
		}

		// 校验用户角色
		role, err := getUserRole(c)
		if err != nil {
			log.Println(err)
			return
		}
		//log.Println(role)
		switch role {
		case model2.UserRoleBanned:
			c.JSON(http.StatusUnauthorized, model.RespError("拒绝访问，用户已被封禁", nil))
			c.Abort()
			return
		case model2.UserRoleAdmin:
			break
		case model2.UserRoleRoot:
			break
		default:
			c.JSON(http.StatusUnauthorized, model.RespError("拒绝访问，用户权限不足", nil))
			c.Abort()
			break
		}

		// 自动刷新token
		err = tokenAutoRefresh(c)
		if err != nil {
			log.Println(err)
			return
		}

		// 放行
		c.Next()
	}
}

func TokenAuthRoot() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 验证token
		err := tokenVerify(c)
		if err != nil {
			log.Println(err)
			return
		}

		// 校验用户角色
		role, err := getUserRole(c)
		if err != nil {
			log.Println(err)
			return
		}
		switch role {
		case model2.UserRoleBanned:
			c.JSON(http.StatusUnauthorized, model.RespError("拒绝访问，用户已被封禁", nil))
			c.Abort()
			return
		case model2.UserRoleRoot:
			break
		default:
			c.JSON(http.StatusUnauthorized, model.RespError("拒绝访问，用户权限不足", nil))
			c.Abort()
			break
		}

		// 自动刷新token
		err = tokenAutoRefresh(c)
		if err != nil {
			log.Println(err)
			return
		}

		// 放行
		c.Next()
	}
}

func tokenAutoRefresh(c *gin.Context) error {
	config := conf.Conf.Token
	exp, err := utils.GetTokenExpire(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, model.RespError("token无效，获取用户信息失败", nil))
		c.Abort()
		return err
	}

	// 计算token剩余时间
	timeLeft := exp - uint64(time.Now().Unix())
	//log.Println(timeLeft, config.Refresh)
	if timeLeft > config.Refresh {
		return nil
	}

	// 获取用户id
	uid, err := utils.GetTokenUid(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, model.RespError("token无效，获取用户信息失败", nil))
		c.Abort()
		return err
	}

	// 生成新token
	token, err := utils.GenerateToken(uid)
	if err != nil {
		c.JSON(http.StatusUnauthorized, model.RespError("token刷新失败", nil))
		c.Abort()
		return err
	}

	c.JSON(http.StatusUnauthorized, model.RespRetry("token已刷新，请重新发送请求", token))
	c.Abort()
	return nil
}

func getUserRole(c *gin.Context) (model2.UserRole, error) {
	// 获取用户id
	uid, err := utils.GetTokenUid(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, model.RespError("token无效，获取用户信息失败", nil))
		c.Abort()
		return 0, err
	}

	// 获取用户信息
	user, err := user.SelectById(uid)
	if err != nil {
		c.JSON(http.StatusUnauthorized, model.RespError("token无效，获取用户信息失败", nil))
		c.Abort()
		return 0, err
	}

	return user.Role, nil
}

func tokenVerify(c *gin.Context) error {
	err := utils.VerifyToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, model.RespError("用户未登录或token过期，请重新登录", nil))
		c.Abort()
		return err
	}

	return nil
}
