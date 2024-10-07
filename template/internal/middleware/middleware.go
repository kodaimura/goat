package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"goat/config"
	"goat/internal/core/jwt"
)


func BasicAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cf := config.GetConfig()

		user, pass, ok := c.Request.BasicAuth()
		if !ok || user != cf.BasicAuthUser || pass != cf.BasicAuthPass {
			c.Header("WWW-Authenticate", "Basic realm=Authorization Required")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}


func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := jwt.Auth(c); err != nil {
			c.Redirect(303, "/login")
			c.Abort()
			return
		}
		c.Next()
	}
}


func JwtAuthApiMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := jwt.Auth(c); err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}