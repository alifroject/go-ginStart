package main

import (
	"gin-quickstart/config"
	"gin-quickstart/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	config.ConnectDatabase()

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173", // React admin panel
			"http://localhost:8081", // Flutter web / testing origin
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	//routes
	routes.RegisterRoutes(router)
	routes.RegisterAuthRoutes(router)
	routes.ProductRoutes(router)

	router.Run(":8080")
}
