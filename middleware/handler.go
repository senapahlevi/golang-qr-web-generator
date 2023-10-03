package middleware

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Claims struct {
	UserID uint `json:"id"`
	jwt.StandardClaims
}

var secret_key = "tokentoken"

func Authentication(c *gin.Context) *Claims {
	userAuth := c.MustGet("user").(*Claims)
	return userAuth
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization header"})
			c.Abort()
			return
		}
		fmt.Println(authHeader, "hello authHeader")
		tokenString := authHeader[len("Bearer "):]
		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
			}
			fmt.Println(token, "hello error: wkwkwk")

			return []byte(secret_key), nil
		})
		fmt.Println(tokenString, "hello tokenString")
		fmt.Println(token, "hello token")

		if err != nil {
			//if expired will showing these message
			c.JSON(http.StatusUnauthorized, gin.H{"status": err.Error()})
			c.Abort()
			return
		}
		claims, ok := token.Claims.(*Claims)
		if !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}
		fmt.Println(claims, "hello claims")
		fmt.Println(ok, "hello ok")

		c.Set("user", claims)
		c.Next()
	}
}
