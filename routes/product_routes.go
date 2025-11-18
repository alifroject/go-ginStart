package routes

import (
    "github.com/gin-gonic/gin"
    "gin-quickstart/controllers"
    "gin-quickstart/middleware"
)

func ProductRoutes(router *gin.Engine) {
    // Group for admin-only routes
    adminProducts := router.Group("/products")
    adminProducts.Use(middleware.AdminOnly()) // Middleware runs first
    {
        adminProducts.POST("/", controllers.CreateProduct)
        adminProducts.PUT("/:id", controllers.UpdateProduct)
        adminProducts.DELETE("/:id", controllers.DeleteProduct)
    }

    // Group for authenticated users (can read products)
    userProducts := router.Group("/products")
    userProducts.Use(middleware.AuthMiddleware()) // Authenticated users only
    {
        userProducts.GET("/", controllers.GetProducts)
        userProducts.GET("/:id", controllers.GetProductByID)
    }
}
