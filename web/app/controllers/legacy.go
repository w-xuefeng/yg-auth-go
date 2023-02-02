package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/w-xuefeng/yg-auth-go/web/app/interfaces"
	"github.com/w-xuefeng/yg-auth-go/web/app/utils"
)

func Auth(c *gin.Context) {
	var body interfaces.AuthBody
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusOK, utils.LegacyJsonFail(err))
		return
	}
	c.JSON(http.StatusOK, utils.LegacyJsonOk(body))
}
