package router

import (
  "github.com/gin-gonic/gin"
  controller "github.com/w-xuefeng/yg-auth-go/src/controllers"
)

func RegisterRouter(r *gin.Engine) {
  r.POST("/auth", controller.Auth)
  r.GET("/welcome", controller.Welcome)
}
