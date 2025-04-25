package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"yorch-devs/bookstore-golang-postgres/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMidleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		// Validates if the auth token is present
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid token"})
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		// Validates if the auth token is valid (secret_key, exp date)
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if ok {
				return utils.JwtSecret, nil
			}

			return nil, jwt.ErrSignatureInvalid
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "The token expired or is invalid"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong when decoding the token"})
			return
		}

		c.Set("user_id", claims["user_id"].(string))

		fmt.Println(claims)
		c.Next()
	}
}
