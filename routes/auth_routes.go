package routes

import (
	"gin-quickstart/controllers"
	"gin-quickstart/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.Engine) {
	router.POST("/signup", controllers.SignUp)
	router.POST("/login", controllers.Login)
	router.GET("/me", middleware.AuthMiddleware(), controllers.GetMe)
}
