package main

import (
    "github.com/gin-gonic/gin"
    "gin-quickstart/config"
    "gin-quickstart/controllers"
)

func main() {
    r := gin.Default()
    config.ConnectDatabase()

    r.GET("/ping", controllers.Ping)
    r.GET("/users", controllers.GetUsers)
    r.POST("/users", controllers.CreateUser)

    r.Run(":8080")
}
