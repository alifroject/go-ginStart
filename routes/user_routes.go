package routes

import (
	"gin-quickstart/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine) {
	userGroup := router.Group("/users")
	{
		userGroup.GET("/", controllers.GetUsers)     
		userGroup.POST("/", controllers.CreateUser)  
	}
}
