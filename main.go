package main

import (
	"gin-quickstart/config"
	"gin-quickstart/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	router := gin.Default()


    //routes
	routes.RegisterRoutes(router)
	routes.RegisterAuthRoutes(router) 


	router.Run(":8080")
}
