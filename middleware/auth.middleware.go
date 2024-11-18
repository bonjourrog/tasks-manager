package middleware

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ValidateToken(c *gin.Context) {
	var (
		key = []byte(os.Getenv("SECRETE_KEY"))
	)
	tokenString := c.Request.Header.Get("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "token is required",
			"data":    nil,
			"error":   true,
		})
		c.Abort()
		return
	}
	t, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			c.Abort()
			return nil, http.ErrNotSupported
		}
		return key, nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
			"data":    nil,
			"error":   true,
		})
		c.Abort()
		return
	}

	// Validación del token y sus claims
	if claims, ok := t.Claims.(jwt.MapClaims); ok {
		// Verificar que el claim "exp" exista y sea un float64
		exp, ok := claims["exp"].(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or missing 'exp' claim"})
			c.Abort()
			return
		}
		// Verificar si el token ha expirado
		if time.Unix(int64(exp), 0).Before(time.Now()) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
			c.Abort()
			return
		}
	} else {
		// Error al convertir los claims
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		c.Abort()
		return
	}
	c.Next()
}
