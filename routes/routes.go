package routes

import (
	"gin-quickstart/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
    router.GET("/ping", controllers.Ping)
}
