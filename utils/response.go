package utils

import "github.com/gin-gonic/gin"

func RespondJSON(c *gin.Context, code int, payload interface{}) {
	c.JSON(code, payload)
}

func RespondError(c *gin.Context, code int, message interface{}) {
	c.JSON(code, gin.H{"error": message})
}