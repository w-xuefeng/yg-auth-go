package controller

import (
	"net/http"

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
	if body.Stuid != "" {
		c.JSON(http.StatusOK, utils.LegacyJsonOk(body.Stuid))
		return
	}
	if body.Token != "" {
		c.JSON(http.StatusOK, utils.LegacyJsonOk(body.Token))
		return
	}
	if body.Regcode != "" {
		c.JSON(http.StatusOK, utils.LegacyJsonOk(body.Regcode))
		return
	}
	c.JSON(http.StatusOK, utils.LegacyJsonFail(config.ResMsgBodyMissing))
}
