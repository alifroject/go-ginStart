package routes

import (
	"gin-quickstart/controllers"
	"gin-quickstart/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.Engine) {
	router.POST("/signup", controllers.SignUp)
	router.POST("/login", controllers.Login)
	router.POST("/admin/login", controllers.AdminLogin)
	router.GET("/me", middleware.AuthMiddleware(), controllers.GetMe)

	 // Protected admin-only routes
    admin := router.Group("/admin")
    admin.Use(middleware.AuthMiddleware(), middleware.AdminOnly())
    {
        admin.GET("/me", controllers.GetMe)
        
    }
}
