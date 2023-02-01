package controller

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

func Welcome(c *gin.Context) {
  c.String(http.StatusOK, "Welcome YG-Auth-Go Server")
}
