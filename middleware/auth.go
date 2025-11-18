package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var JWT_SECRET = []byte(os.Getenv("JWT_SECRET"))

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        var tokenString string

        // 1. Try cookie (Admin Web)
        cookie, err := c.Cookie("token")
        if err == nil {
            tokenString = cookie
        }

        // 2. Try header (Flutter)
        if tokenString == "" {
            authHeader := c.GetHeader("Authorization")
            if authHeader != "" {
                parts := strings.Split(authHeader, " ")
                if len(parts) == 2 && parts[0] == "Bearer" {
                    tokenString = parts[1]
                }
            }
        }

        if tokenString == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
            return
        }

        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return JWT_SECRET, nil
        })

        if err != nil || !token.Valid {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
            return
        }

        if claims, ok := token.Claims.(jwt.MapClaims); ok {
            c.Set("user", claims)
            c.Set("role", claims["role"])
        }

        c.Next()
    }
}


// admin
func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || role != "admin" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Forbidden: Admins only"})
			return
		}
		c.Next()
	}
}
