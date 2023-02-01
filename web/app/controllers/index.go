package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/w-xuefeng/yg-auth-go/web/app/config"
	"github.com/w-xuefeng/yg-auth-go/web/app/utils"
	"net/http"
)

func Index(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, config.AuthIndex)
}

func Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, utils.UniJson("welcome", true, "success", 200))
}
