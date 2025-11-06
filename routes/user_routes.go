package routes

import (
	"gin-quickstart/controllers"
	"gin-quickstart/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine) {
	userGroup := router.Group("/user")
	{
		// GET /user - all users, admin only
		userGroup.GET("/", middleware.AuthMiddleware(), middleware.AdminOnly(), controllers.GetUsers)
		//no auth
		userGroup.POST("/", controllers.CreateUser)
	}
}
