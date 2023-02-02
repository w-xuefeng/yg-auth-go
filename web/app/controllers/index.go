package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/w-xuefeng/yg-auth-go/web/app/config"
	"github.com/w-xuefeng/yg-auth-go/web/app/utils"
)

func Index(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, config.AuthIndex)
}

func Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, utils.UniJsonOk("welcome"))
}
