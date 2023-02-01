package router

import (
	"github.com/gin-gonic/gin"
	controller "github.com/w-xuefeng/yg-auth-go/web/app/controllers"
)

func RegisterRouter(r *gin.Engine) {
	r.GET("/", controller.Index)
	r.POST("/auth", controller.Auth)
	r.GET("/welcome", controller.Welcome)
}
