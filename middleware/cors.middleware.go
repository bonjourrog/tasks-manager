package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func CorsMiddleware(c *gin.Context) {
	allowedOrigin := os.Getenv("ALLOWED_ORIGIN")
	if allowedOrigin != "" {
		c.Writer.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
	} else {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	}
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}
	c.Next()
}
