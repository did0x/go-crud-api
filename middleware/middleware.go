package middleware

import (
	"fmt"
	"strings"
	"net/http"
	"github.com/gin-gonic/gin"
	jwt "github.com/dgrijalva/jwt-go"
)

// CommonMiddleware to give default setting to router
func CommonMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		c.Next()
	}
}

// AuthMiddleware to give authentication to router
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if len(tokenString) == 0 {
			c.String(http.StatusUnauthorized, "Missing Authorization Header")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		claims, err := VerifyToken(tokenString)
		if err != nil {
			c.String(http.StatusUnauthorized, fmt.Sprintf("Error verifying JWT token: %s" ,err.Error()))
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		username := claims.(jwt.MapClaims)["username"].(string)
		c.Set("username", username)

		c.Next()
	}
}
