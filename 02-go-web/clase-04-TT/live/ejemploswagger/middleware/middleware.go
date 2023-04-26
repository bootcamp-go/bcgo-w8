package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"log"
)

func MiddlewareUno() gin.HandlerFunc {
	log.Println("Este es el primer middleware")

	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != "123" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "El token no es el mismo"})
			return
		}
		c.Next()
	}
}
