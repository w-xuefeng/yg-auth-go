package controller

import (
	"net/http"

	"github.com/w-xuefeng/yg-auth-go/web/app/services"

	"github.com/gin-gonic/gin"
	"github.com/w-xuefeng/yg-auth-go/web/app/config"
	"github.com/w-xuefeng/yg-auth-go/web/app/interfaces"
	"github.com/w-xuefeng/yg-auth-go/web/app/utils"
)

func Auth(c *gin.Context) {
	var body interfaces.AuthBody
	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusOK, utils.LegacyJsonFail(config.ResMsgBodyMissing))
		return
	}
	if body.Stuid != "" && body.Password != "" {
		loginInfo, err := services.Login(body.Stuid, body.Password)
		if err != nil {
			c.JSON(http.StatusOK, utils.LegacyJsonFail(err.Error()))
			return
		}
		c.JSON(http.StatusOK, utils.LegacyJsonOk(loginInfo))
		return
	}
	if body.Token != "" {
		userInfo, err := services.GetUserByToken(body.Token)
		if err != nil {
			c.JSON(http.StatusOK, utils.LegacyJsonFail(err.Error()))
			return
		}
		c.JSON(http.StatusOK, utils.LegacyJsonOk(userInfo))
		return
	}
	if body.Regcode != "" {
		next, err := services.CheckRegCode(body.Regcode)
		if err != nil {
			c.JSON(http.StatusOK, utils.LegacyJsonFail(err.Error()))
			return
		}
		c.JSON(http.StatusOK, interfaces.LegacyResponse[any]{Status: next})
		return
	}
	c.JSON(http.StatusOK, utils.LegacyJsonFail(config.ResMsgBodyMissing))
}
