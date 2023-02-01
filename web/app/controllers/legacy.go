package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/w-xuefeng/yg-auth-go/web/app/utils"
	"net/http"
)

func Auth(c *gin.Context) {
	c.JSON(http.StatusOK, utils.UniJson(gin.H{}, false, "fail", 1400))
}
