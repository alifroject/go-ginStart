package routes

import (
  "github.com/gin-gonic/gin"
  "gin-quickstart/controllers"
)

func RegisterRoutes(router *gin.Engine) {
  router.GET("/ping", controllers.Ping)
}
